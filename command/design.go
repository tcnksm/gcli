package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/tcnksm/gcli/skeleton"
	"github.com/tcnksm/go-gitconfig"
)

const (
	defaultOutputFmt = "%s-design.toml"
)

// DesignCommand is a Command that generates a new cli project
type DesignCommand struct {
	Meta
}

// Run generates a new cli project. It returns exit code
func (c *DesignCommand) Run(args []string) int {

	var (
		output       string
		owner        string
		commands     []*skeleton.Command
		flags        []*skeleton.Flag
		frameworkStr string
	)

	uflag := c.Meta.NewFlagSet("design", c.Help())

	uflag.Var((*CommandFlag)(&commands), "command", "command")
	uflag.Var((*CommandFlag)(&commands), "c", "command (short)")

	uflag.Var((*FlagFlag)(&flags), "flag", "flag")
	uflag.Var((*FlagFlag)(&flags), "f", "flag (short)")

	uflag.StringVar(&frameworkStr, "framework", defaultFrameworkString, "framework")
	uflag.StringVar(&frameworkStr, "F", defaultFrameworkString, "framework (short)")

	uflag.StringVar(&owner, "owner", "", "owner")
	uflag.StringVar(&owner, "o", "", "owner (short)")

	uflag.StringVar(&output, "output", "", "output")
	uflag.StringVar(&output, "O", "", "output (short)")

	if err := uflag.Parse(args); err != nil {
		return 1
	}

	parsedArgs := uflag.Args()
	if len(parsedArgs) != 1 {
		msg := fmt.Sprintf("Invalid arguments: usage gcli design [option] NAME")
		c.UI.Error(msg)
		return 1
	}

	name := parsedArgs[0]

	// If output file name is not provided use default one
	if len(output) == 0 {
		output = fmt.Sprintf(defaultOutputFmt, name)
	}

	if _, err := os.Stat(output); !os.IsNotExist(err) {
		msg := fmt.Sprintf("Cannot create design file %s: file exists", output)
		c.UI.Error(msg)
		return 1
	}

	outputFile, err := os.Create(output)
	if err != nil {
		msg := fmt.Sprintf("Cannot create design file %s: %s", output, err)
		c.UI.Error(msg)
		return 1
	}

	if owner == "" {
		owner, err = gitconfig.GithubUser()
		if err != nil {
			owner, _ = gitconfig.Username()
		}
	}

	// If no commands are specified, set emply value so that
	// user can understand how to write
	if len(commands) < 1 && len(flags) < 1 {
		commands = []*skeleton.Command{
			&skeleton.Command{
				Name: "",
			},
		}
	}

	// Define Executable
	executable := &skeleton.Executable{
		Name:         name,
		Owner:        owner,
		Commands:     commands,
		Flags:        flags,
		Version:      skeleton.DefaultVersion,
		Description:  skeleton.DefaultDescription,
		FrameworkStr: frameworkStr,
	}

	if err := toml.NewEncoder(outputFile).Encode(executable); err != nil {
		msg := fmt.Sprintf("Failed to generate design file: %s", err)
		c.UI.Error(msg)
		return 1
	}

	c.UI.Info(fmt.Sprintf("====> Successfully generated %s", output))
	return ExitCodeOK
}

// Synopsis is a one-line, short synopsis of the command.
func (c *DesignCommand) Synopsis() string {
	return "Generate project design template"
}

// Help is a long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (c *DesignCommand) Help() string {
	helpText := `
Generate project design template (as toml file). You can pass that file to 'gcli apply'
command and generate CLI tool based on template file. You can define what command
and what flag you need on that file.

Usage:

  gcli design [option] NAME


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

  -framework=name, -F         Cli framework name. By default, gcli use "urfave/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

  -owner=name, -o             Command owner (author) name. This value is also used for
                              import path name. By default, owner name is extracted from
                              ~/.gitconfig variable.


  -output, -O                 Change output file name. By default, gcli use "NAME-design.toml"
`
	return strings.TrimSpace(helpText)
}
