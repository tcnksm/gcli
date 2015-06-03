package command

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tcnksm/cli-init/skeleton"
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

func (c *NewCommand) Run(args []string) int {

	var commands []skeleton.Command
	var flags []skeleton.Flag
	var frameworkStr string
	var owner string
	var skipTest bool

	uflag := flag.NewFlagSet("new", flag.ContinueOnError)
	uflag.Usage = func() { c.Ui.Error(c.Help()) }

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

	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	uflag.SetOutput(errW)

	go func() {
		for errScanner.Scan() {
			c.Ui.Error(errScanner.Text())
		}
	}()

	if err := uflag.Parse(args); err != nil {
		return 1
	}

	parsedArgs := uflag.Args()
	if len(parsedArgs) != 1 {
		msg := fmt.Sprintf("invalid arguments")
		c.Ui.Error(msg)
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
		c.Ui.Error(msg)
		return 1
	}

	framework, err := skeleton.Framework(frameworkStr)
	if err != nil {
		return 1
	}

	// Use .gitconfig value.
	if owner == "" {
		owner, err = gitconfig.GithubUser()
		if err != nil {
			owner, err = gitconfig.Username()
			if err != nil {
				msg := "Cannot retrieve owner name\n" +
					"Owener name is retrieved from `~/.gitcofig` file.\n" +
					"Please set one via -owner option or `~/.gitconfig` file."
				c.Ui.Error(msg)
				return 1
			}
		}
	}

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
	}

	// Create project directory
	errCh := skeleton.Generate()

	gotErr := false
	for err := range errCh {
		gotErr = true
		c.Ui.Error(err.Error())
	}

	// Return non zero var when at least one
	// error was happened while executing skeleton.Generate().
	// Run all templating and show all error.
	if gotErr {
		c.Ui.Error(fmt.Sprintf("Failed to generate: %s", name))
		return 1
	}

	c.Ui.Info(fmt.Sprintf("====> Successfuly generated: %s", name))
	return 0
}

func (c *NewCommand) Synopsis() string {
	return "Generate new cli project"
}

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
