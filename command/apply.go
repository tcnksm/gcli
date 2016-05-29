package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/tcnksm/gcli/skeleton"
)

// ApplyCommand is a Command that generates a new cli project
type ApplyCommand struct {
	Meta
}

// Run generates a new cli project. It returns exit code
func (c *ApplyCommand) Run(args []string) int {

	var (
		frameworkStr string
		owner        string
		name         string
		staticDir    string
		vcsHost      string
		current      bool
		skipTest     bool
		verbose      bool
	)

	uflag := c.Meta.NewFlagSet("apply", c.Help())

	uflag.StringVar(&frameworkStr, "framework", "", "framework")
	uflag.StringVar(&frameworkStr, "F", "", "framework (short)")

	uflag.StringVar(&staticDir, "static-dir", "", "")

	uflag.StringVar(&vcsHost, "vcs", DefaultVCSHost, "")

	uflag.BoolVar(&current, "current", false, "current")
	uflag.BoolVar(&current, "C", false, "current")

	uflag.BoolVar(&skipTest, "skip-test", false, "skip-test")
	uflag.BoolVar(&skipTest, "T", false, "skip-test (short)")

	uflag.BoolVar(&verbose, "verbose", false, "verbose")
	uflag.BoolVar(&verbose, "V", false, "verbose (short)")

	// These flags are supposed only to use in test
	uflag.StringVar(&owner, "owner", "", "owner (Should only for test)")
	uflag.StringVar(&name, "name", "", "name (Should only for test)")

	if err := uflag.Parse(args); err != nil {
		return 1
	}

	parsedArgs := uflag.Args()
	if len(parsedArgs) != 1 {
		c.UI.Error("Invalid argument: Usage glic apply [options] FILE")
		return 1
	}

	designFile := parsedArgs[0]
	c.UI.Info(fmt.Sprintf(
		"Use design template %q for generating new cli project", designFile))

	// Check file is exist or not
	if _, err := os.Stat(designFile); os.IsNotExist(err) {
		c.UI.Error(fmt.Sprintf(
			"Design file does not exsit"))
		return 1
	}

	// Decode design file as skeleton.Executable
	executable := skeleton.NewExecutable()
	if _, err := toml.DecodeFile(designFile, executable); err != nil {
		c.UI.Error(fmt.Sprintf(
			"Failed to decode design file %q: %s", designFile, err))
		return 1
	}

	if err := executable.Fix(); err != nil {
		c.UI.Error(fmt.Sprintf(
			"Failed to fix input value: %s", err))
		return 1
	}

	// validate executable
	if errs := executable.Validate(); len(errs) > 0 {
		c.UI.Error(fmt.Sprintf(
			"%q is not valid template file. It has %d errors:", designFile, len(errs)))
		for _, err := range errs {
			c.UI.Error(fmt.Sprintf(
				"  * %s", err.Error()))
		}
		return ExitCodeFailed
	}

	// Check option input first and if it's specified use it
	if len(frameworkStr) == 0 {
		if len(executable.FrameworkStr) != 0 {
			// If FrameworStr is specified from design file use it
			frameworkStr = executable.FrameworkStr
		} else {
			frameworkStr = defaultFrameworkString
		}
	}

	// Overwrite VCSHost
	executable.VCSHost = vcsHost

	// Overwrite output and executable name
	output := executable.Name
	if len(name) != 0 {
		executable.Name = name
		output = name
	}

	// Overwrite owner name
	if len(owner) != 0 {
		executable.Owner = owner
	}

	currentDir, err := os.Getwd()
	if err != nil {
		c.UI.Error(fmt.Sprintf(
			"Failed to get current directroy: %s", err))
		return ExitCodeFailed
	}

	gopath := os.Getenv(EnvGoPath)
	if gopath == "" {
		c.UI.Error(fmt.Sprintf(
			"Failed to read GOPATH: it should not be empty"))
		return ExitCodeFailed
	}
	idealDir := filepath.Join(gopath, "src", executable.VCSHost, executable.Owner)

	if currentDir != idealDir && !current {
		c.UI.Output("")
		c.UI.Output(fmt.Sprintf("====> WARNING: You are not in the directory gcli expects."))
		c.UI.Output(fmt.Sprintf("      The codes will be generated be in $GOPATH/src/%s/%s.",
			executable.VCSHost, executable.Owner))
		c.UI.Output(fmt.Sprintf("      Not in the current directory. This is because the output"))
		c.UI.Output(fmt.Sprintf("      codes use import path based on that path."))
		c.UI.Output("")
		output = filepath.Join(idealDir, executable.Name)
	}

	if _, err := os.Stat(output); !os.IsNotExist(err) {
		msg := fmt.Sprintf("Cannot create directory %s: file exists", output)
		c.UI.Error(msg)
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

	framework, err := skeleton.FrameworkByName(frameworkStr)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Failed to generate %q: %s", executable.Name, err.Error()))
		return 1
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
			c.UI.Info(fmt.Sprintf("====> Successfully generated %s", executable.Name))
			return ExitCodeOK
		}
	}
}

// Synopsis is a one-line, short synopsis of the command.
func (c *ApplyCommand) Synopsis() string {
	return "Apply design template file for generating cli project"
}

// Help is a long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (c *ApplyCommand) Help() string {
	helpText := `
Apply design template file for generating cli project. You can generate
design template file via 'gcli design' command. If framework name is not
specified gcli use urfave/cli. You can set framework name via '-F'
option. To check cli framework you can use, run 'gcli list'. 

Usage:

  gcli apply [option] FILE


Options:

   -framework=name, -F        Cli framework name. By default, gcli use "urfave/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

   -vcs=name                  Version Control Host name. By default, gcli use 'github.com'.

   -skip-test, -T             Skip generating *_test.go file. By default, gcli generates
                              test file If you specify this flag, gcli will not generate
                              test files.
`
	return strings.TrimSpace(helpText)
}
