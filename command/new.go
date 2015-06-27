package command

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tcnksm/gcli/skeleton"
	"github.com/tcnksm/go-gitconfig"
)

const (
	// defaultFrameworkString is default cli framework name
	defaultFrameworkString = "codegangsta_cli"

	// defaultVersion is default appliaction version
	defaultVersion = "0.1.0"

	// defaultVersion is default application description
	defaultDescription = ""
)

// NewCommand is a Command that generates a new cli project
type NewCommand struct {
	Meta
}

// Run generates a new cli project. It returns exit code
func (c *NewCommand) Run(args []string) int {

	var (
		commands     []skeleton.Command
		flags        []skeleton.Flag
		frameworkStr string
		owner        string
		skipTest     bool
		verbose      bool
	)

	uflag := flag.NewFlagSet("new", flag.ContinueOnError)
	uflag.Usage = func() { c.UI.Error(c.Help()) }

	uflag.Var((*CommandFlag)(&commands), "command", "command")
	uflag.Var((*CommandFlag)(&commands), "c", "command (short)")

	uflag.Var((*FlagFlag)(&flags), "flag", "flag")
	uflag.Var((*FlagFlag)(&flags), "f", "flag (short)")

	uflag.StringVar(&frameworkStr, "framework", defaultFrameworkString, "framework")
	uflag.StringVar(&frameworkStr, "F", defaultFrameworkString, "framework (short)")

	uflag.StringVar(&owner, "owner", "", "owner")
	uflag.StringVar(&owner, "o", "", "owner (short)")

	uflag.BoolVar(&skipTest, "skip-test", false, "skip-test")
	uflag.BoolVar(&skipTest, "T", false, "skip-test (short)")

	uflag.BoolVar(&verbose, "verbose", false, "verbose")
	uflag.BoolVar(&verbose, "V", false, "verbose (short)")

	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	uflag.SetOutput(errW)

	go func() {
		for errScanner.Scan() {
			c.UI.Error(errScanner.Text())
		}
	}()

	if err := uflag.Parse(args); err != nil {
		return 1
	}

	parsedArgs := uflag.Args()
	if len(parsedArgs) != 1 {
		msg := fmt.Sprintf("Invalid arguments: %s", strings.Join(parsedArgs, " "))
		c.UI.Error(msg)
		return 1
	}

	name := parsedArgs[0]

	// TODO, should be configurable
	// or chagne direcotry to GOPATH/github.com/owner/output
	// Some gcli template assume command is executed
	// from GOPATH/github.com/owner
	output := name
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		msg := fmt.Sprintf("Cannot create directory %s: file exists", output)
		c.UI.Error(msg)
		return 1
	}

	framework, err := skeleton.FrameworkByName(frameworkStr)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Failed to generate %q: %s", name, err.Error()))
		return 1
	}

	// Use .gitconfig value.
	if owner == "" {
		owner, err = gitconfig.GithubUser()
		if err != nil {
			owner, err = gitconfig.Username()
			if err != nil {
				msg := "Cannot find owner name\n" +
					"By default, owener name is retrieved from `~/.gitcofig` file.\n" +
					"Please set one via -owner option or `~/.gitconfig` file."
				c.UI.Error(msg)
				return 1
			}
		}
	}

	// outCh receives info output from skeleton.Generate
	// and show it in UI output
	outCh := make(chan string)
	go func() {
		for out := range outCh {
			c.UI.Output("  " + out)
		}
	}()

	// errCh receives error from skeleton.Generate
	// and show it in UI error output
	gotErr := false
	err2Ch := make(chan error)
	go func() {
		for err := range err2Ch {
			c.UI.Error("  " + err.Error())
		}
	}()

	executable := &skeleton.Executable{
		Name:        name,
		Owner:       owner,
		Commands:    commands,
		Flags:       flags,
		Version:     defaultVersion,
		Description: defaultDescription,
	}

	skeleton := &skeleton.Skeleton{
		Path:       output,
		Framework:  framework,
		SkipTest:   skipTest,
		Executable: executable,
		OutCh:      outCh,
		ErrCh:      err2Ch,
		Verbose:    verbose,
		LogWriter:  os.Stdout,
	}

	// Create project directory
	doneCh := skeleton.Generate()
	<-doneCh

	// Return non zero var when at least one
	// error was happened while executing skeleton.Generate().
	// Run all templating and show all error.
	if gotErr {
		c.UI.Error(fmt.Sprintf("Failed to generate %q", name))
		return 1
	}

	c.UI.Info(fmt.Sprintf("====> Successfuly generated: %s", name))
	return 0
}

// Synopsis is a one-line, short synopsis of the command.
func (c *NewCommand) Synopsis() string {
	return "Generate new cli project"
}

// Help is a long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (c *NewCommand) Help() string {
	helpText := `
Usage: gcli new [option] NAME

  Generate new cli skeleton project. At least, you must provide executable
  name. You can select cli package and set commands via command line option.
  See more about that on Options section. By default, gcli use codegangsta/cli.
  To check cli framework you can use, run 'gcli list'. 

Options:

  -command=name, -c           Command name which you want to add.
                              This is valid only when cli pacakge support commands.
                              This can be specified multiple times. Synopsis can be
                              set after ":". Namely, you can specify command by 
                              -command=NAME:SYNOPSYS. Only NAME is required.
                              You can set multiple variables at same time with ","
                              separator.

  -flag=name, -f              Global flag option name which you want to add.
                              This can be specified multiple times. By default, flag type
                              is string and its description is empty. You can set them,
                              with ":" separator. Namaly, you can specify flag by
                              -flag=NAME:TYPE:DESCIRPTION. Order must be flow  this and
                              TYPE must be string, bool or int. Only NAME is required.
                              You can set multiple variables at same time with ","
                              separator.

   -framework=name, -F        Cli framework name. By default, gcli use "codegangsta/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

   -owner=name, -o            Command owner (author) name. This value is also used for
                              import path name. By default, owner name is extracted from
                              ~/.gitconfig variable.

   -skip-test, -T             Skip generating *_test.go file. By default, gcli generates
                              test file If you specify this flag, gcli will not generate
                              test files.

Examples:

   This example shows creating todo command application skeleton
   which has 'add' and 'delete' command by using mitchellh/cli package.

   $ gcli new -command=add:"Add new task" -commnad=delete:"delete task" todo
`
	return strings.TrimSpace(helpText)
}
