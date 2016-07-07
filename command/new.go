package command

import (
	"fmt"
	"os"
	"path/filepath"
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
		commands     []*skeleton.Command
		flags        []*skeleton.Flag
		frameworkStr string
		owner        string
		staticDir    string
		vcsHost      string
		current      bool
		skipTest     bool
		verbose      bool
	)

	uflag := c.Meta.NewFlagSet("new", c.Help())

	uflag.Var((*CommandFlag)(&commands), "command", "command")
	uflag.Var((*CommandFlag)(&commands), "c", "command (short)")

	uflag.Var((*FlagFlag)(&flags), "flag", "flag")
	uflag.Var((*FlagFlag)(&flags), "f", "flag (short)")

	uflag.StringVar(&frameworkStr, "framework", defaultFrameworkString, "framework")
	uflag.StringVar(&frameworkStr, "F", defaultFrameworkString, "framework (short)")

	uflag.StringVar(&owner, "owner", "", "owner")
	uflag.StringVar(&owner, "o", "", "owner (short)")

	uflag.StringVar(&staticDir, "static-dir", "", "")

	uflag.StringVar(&vcsHost, "vcs", DefaultVCSHost, "")

	uflag.BoolVar(&current, "current", false, "current")
	uflag.BoolVar(&current, "C", false, "current")

	uflag.BoolVar(&skipTest, "skip-test", false, "skip-test")
	uflag.BoolVar(&skipTest, "T", false, "skip-test (short)")

	uflag.BoolVar(&verbose, "verbose", false, "verbose")
	uflag.BoolVar(&verbose, "V", false, "verbose (short)")

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

	// If owner is not provided, use .gitconfig value.
	if owner == "" {
		var err error
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

	currentDir, err := os.Getwd()
	if err != nil {
		c.UI.Error(fmt.Sprintf(
			"Failed to get current directroy: %s", err))
		return ExitCodeFailed
	}

	gopaths := filepath.SplitList(os.Getenv(EnvGoPath))
	gopath := ""
	if len(gopaths) == 0 {
		c.UI.Error(fmt.Sprintf(
			"Failed to read GOPATH: it should not be empty"))
		return ExitCodeFailed
	} else {
		for _, path := range gopaths {
			absPath, err := filepath.Abs(path)
			if err != nil {
				c.UI.Error(fmt.Sprintf(
					"Cannot parse GOPATH"))
				continue
			}
			if strings.HasPrefix(currentDir, absPath) {
				gopath = absPath
				break
			}
		}
	}
	if gopath == "" {
		c.UI.Output("")
		c.UI.Output(fmt.Sprintf("===> WARNING: You are not in the directories defined in $GOPATH."))
		c.UI.Output(fmt.Sprintf("     Uses first location in $GOPATH."))
		c.UI.Output("")
		gopath, err = filepath.Abs(gopaths[0])
		if err != nil {
			c.UI.Error(fmt.Sprintf("Cannot parse GOPATH"))
			return ExitCodeFailed
		}
	}
	
	idealDir := filepath.Join(gopath, "src", vcsHost, owner)

	output := name
	if currentDir != idealDir && !current {
		c.UI.Output("")
		c.UI.Output(fmt.Sprintf("====> WARNING: You are not in the directory gcli expects."))
		c.UI.Output(fmt.Sprintf("      The codes will be generated be in $GOPATH/src/%s/%s.", vcsHost, owner))
		c.UI.Output(fmt.Sprintf("      Not in the current directory. This is because the output"))
		c.UI.Output(fmt.Sprintf("      codes use import path based on that path."))
		c.UI.Output("")
		output = filepath.Join(idealDir, name)
	}

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

	if staticDir == "" {
		localDir, err := c.LocalDir()
		if err != nil {
			c.UI.Error(err.Error())
			return ExitCodeFailed
		}
		staticDir = filepath.Join(localDir, DefaultLocalStaticDir)
	}

	// Define Executable
	executable := &skeleton.Executable{
		Name:        name,
		Owner:       owner,
		VCSHost:     vcsHost,
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
		StaticDir:  staticDir,
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
Generate new cli skeleton project. At least, you must provide executable
name. You can select cli package and set commands via command line option.
See more about that on Options section. By default, gcli use codegangsta/cli.
To check cli framework you can use, run 'gcli list'. 

Usage:

    gcli new [option] NAME

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

   -vcs=name                  Version Control Host name. By default, gcli use 'github.com'.

   -skip-test, -T             Skip generating *_test.go file. By default, gcli generates
                              test file If you specify this flag, gcli will not generate
                              test files.

Examples:

To create todo command application skeleton which has 'add' and 'delete' command,

   $ gcli new -command=add:"Add new task" -command=delete:"delete task" todo
`
	return strings.TrimSpace(helpText)
}
