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

	// Define Executable
	executable := &skeleton.Executable{
		Name:        name,
		Owner:       owner,
		Commands:    commands,
		Flags:       flags,
		Version:     skeleton.DefaultVersion,
		Description: skeleton.DefaultDescription,
	}

	// Channels to receive artifact path (result) and error
	artifactCh, errCh := make(chan string), make(chan error)

	// Define Skeleton
	skeleton := &skeleton.Skeleton{
		Path:       output,
		Framework:  framework,
		SkipTest:   skipTest,
		Executable: executable,
		ArtifactCh: artifactCh,
		ErrCh:      errCh,
		Verbose:    verbose,
		LogWriter:  os.Stdout,
	}

	// Create project directory
	doneCh := skeleton.Generate()

	for {
		select {
		case artifact := <-artifactCh:
			c.UI.Output(fmt.Sprintf("  Created %s", artifact))
		case err := <-errCh:
			c.UI.Error(fmt.Sprintf("Failed to generate %s: %s", output, err.Error()))

			// If some file are created before error happend
			// Should be cleanuped
			if _, err := os.Stat(output); !os.IsNotExist(err) {
				c.UI.Output(fmt.Sprintf("Cleanup %s", output))
				os.RemoveAll(output)
			}
			return ExitCodeFailed
		case <-doneCh:
			c.UI.Info(fmt.Sprintf("====> Successfully generated %s", name))
			return ExitCodeOK
		}
	}
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
