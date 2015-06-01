package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/tcnksm/cli-init/skeleton"
)

const (
	// defaultTypeString is default flag type
	defaultTypeString = "string"

	// defaultFrameworkString is default cli framework name
	defaultFrameworkString = "codegangsta_cli"
)

// NewCommand is a Command that generates a new cli project
type NewCommand struct {
	Meta
}

func (c *NewCommand) Run(args []string) int {

	var frameworkStr string
	var commands []skeleton.Command
	var flags []skeleton.Flag

	uflag := flag.NewFlagSet("new", flag.ContinueOnError)
	uflag.Var((*CommandFlag)(&commands), "command", "command")
	uflag.Var((*FlagFlag)(&flags), "flag", "flag")

	uflag.StringVar(&frameworkStr, "framework", defaultFrameworkString, "framework")

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

	framework, err := skeleton.Framework(frameworkStr)
	if err != nil {
		return 1
	}

	executable := &skeleton.Executable{
		Name:        name,
		Owner:       "tcnksm",
		Commands:    commands,
		Flags:       flags,
		Version:     "0.1.0",
		Description: "todo is CLI todo manager",
	}

	skeleton := &skeleton.Skeleton{
		Path:       name,
		Framework:  framework,
		WithTest:   true,
		Executable: executable,
	}

	// Create project directory
	errCh := skeleton.Generate()

	gotErr := false
	for err := range errCh {
		gotErr = true
		c.Ui.Error(err.Error())
	}

	// Return non zero var when more than one
	// error is happed while executing skeleton.Generat().
	// Run all templating and show all error.
	if gotErr {
		c.Ui.Error(fmt.Sprintf("Failed to generate: %s", name))
		return 1
	}

	c.Ui.Info(fmt.Sprintf("====> Successfuly generated: %s", name))
	return 0
}

// FlagFlag implements the flag.Value interface and allows multiple
// calls to the same variable to append a list. It parses string and set them
// as skeleton.Flag.
type FlagFlag []skeleton.Flag

func (f *FlagFlag) String() string {
	return ""
}

func (f *FlagFlag) Set(v string) error {

	parsed := strings.Split(v, ":")
	if len(parsed) > 3 {
		return fmt.Errorf("flag must be specified by NAME:TYPE:DESCRIPTION format")
	}

	name := parsed[0]
	typeString := defaultTypeString
	desc := ""

	if len(parsed) > 1 {
		typeString = parsed[1]
	}

	if len(parsed) > 2 {
		desc = parsed[2]
	}

	flag := skeleton.Flag{
		LongName:    name,
		TypeString:  typeString,
		Description: desc,
	}

	// Fix inputs string for using main processing
	if err := flag.Fix(); err != nil {
		return err
	}

	*f = append(*f, flag)
	return nil
}

// CommandFlag implements the flag.Value interface and allows multiple
// calls to the same variable to append a list. It parses string and set them
// as skeleton.Command.
type CommandFlag []skeleton.Command

func (c *CommandFlag) String() string {
	return ""
}

func (c *CommandFlag) Set(v string) error {
	parsed := strings.Split(v, ":")
	if len(parsed) > 2 {
		return fmt.Errorf("command flag must be specified by NAME:SYNOPSIS format")
	}

	name := parsed[0]

	// synopsis is optional
	synopsis := ""
	if len(parsed) == 2 {
		synopsis = parsed[1]
	}

	*c = append(*c, skeleton.Command{
		Name:     name,
		Synopsis: synopsis,
	})

	return nil
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

  -command=name               Command name which you want to add to executable.
                              This is valid only when cli pacakge support commands.
                              This can be specified multiple times. Synopsis can be
                              set after ":". Namely, you can specify command by 
                              -command=NAME:SYNOPSYS. Only NAME is required. 

  -flag=name                  Global flag option name which you want to add to executable.
                              This can be specified multiple times. By default, flag type
                              is string and its description is empty. You can set them,
                              with ":" separator. Namaly, you can specify flag by
                              -flag=NAME:TYPE:DESCIRPTION. Order must be flow  this and
                              TYPE must be string, bool or int. Only NAME is required.

   -framework=name            Cli framework name. By default, gcli use codegangsta/cli
                              To check cli framework you can use, run 'gcli list'.

   -owner=name                Command owner (author) name. This value is used import path name.
                              By default, owner name is extracted from gitconfig variable.
                              This is optional.

Examples:

   This example shows creating todo command application skeleton
   which has 'add' and 'delete' command by mitchellh/cli package.

   $ gcli new -command=add:"Add new task" -commnad=delete:"delete task" todo

`
	return strings.TrimSpace(helpText)
}
