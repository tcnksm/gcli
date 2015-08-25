package command

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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
		skipTest     bool
		verbose      bool
		owner        string
		name         string
	)

	uflag := flag.NewFlagSet("apply", flag.ContinueOnError)
	uflag.Usage = func() { c.UI.Error(c.Help()) }

	uflag.StringVar(&frameworkStr, "framework", "", "framework")
	uflag.StringVar(&frameworkStr, "F", "", "framework (short)")

	uflag.BoolVar(&skipTest, "skip-test", false, "skip-test")
	uflag.BoolVar(&skipTest, "T", false, "skip-test (short)")

	uflag.BoolVar(&verbose, "verbose", false, "verbose")
	uflag.BoolVar(&verbose, "V", false, "verbose (short)")

	// These flags are supposed only to use in test
	uflag.StringVar(&owner, "owner", "", "owner (Should only for test)")
	uflag.StringVar(&name, "name", "", "name (Should only for test)")

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
		c.UI.Error("Invalid argument: Usage glic apply [options] FILE")
		return 1
	}

	planFile := parsedArgs[0]
	c.UI.Info(fmt.Sprintf(
		"Use plan file %q for generating new cli project", planFile))

	// Check file is exist or not
	if _, err := os.Stat(planFile); os.IsNotExist(err) {
		c.UI.Error(fmt.Sprintf(
			"Plan file does not exsit"))
		return 1
	}

	// Decode plan file as skeleton.Executable
	executable := skeleton.NewExecutable()
	if _, err := toml.DecodeFile(planFile, executable); err != nil {
		c.UI.Error(fmt.Sprintf(
			"Failed to decode plan file %q: %s", planFile, err))
		return 1
	}

	output := executable.Name
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		msg := fmt.Sprintf("Cannot create directory %s: file exists", output)
		c.UI.Error(msg)
		return 1
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

	fmt.Println(frameworkStr)
	framework, err := skeleton.FrameworkByName(frameworkStr)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Failed to generate %q: %s", executable.Name, err.Error()))
		return 1
	}

	if len(name) != 0 {
		executable.Name = name
		output = name
	}

	if len(owner) != 0 {
		executable.Owner = owner
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
Usage: gcli apply [option] FILE

  Apply design template file for generating cli project. You can generate
  design template file via 'gcli design' command. If framework name is not
  specified gcli use codegangsta/cli. You can set framework name via '-F'
  option. To check cli framework you can use, run 'gcli list'. 

Options:

   -framework=name, -F        Cli framework name. By default, gcli use "codegangsta/cli"
                              To check cli framework you can use, run 'gcli list'.
                              If you set invalid framework, it will be failed.

   -skip-test, -T             Skip generating *_test.go file. By default, gcli generates
                              test file If you specify this flag, gcli will not generate
                              test files.
`
	return strings.TrimSpace(helpText)
}
