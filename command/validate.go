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

// ValidateCommand is a Command that validate template file
type ValidateCommand struct {
	Meta
}

// Run validates template file
func (c *ValidateCommand) Run(args []string) int {

	uflag := flag.NewFlagSet("validate", flag.ContinueOnError)
	uflag.Usage = func() { c.UI.Error(c.Help()) }

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
		c.UI.Error("Invalid argument: Usage glic validate [options] FILE")
		return 1
	}

	designFile := parsedArgs[0]

	// Check file is exist or not
	if _, err := os.Stat(designFile); os.IsNotExist(err) {
		c.UI.Error(fmt.Sprintf(
			"Design file %q does not exsit: %s", designFile, err))
		return 1
	}

	// Decode design file as skeleton.Executable
	executable := skeleton.NewExecutable()
	if _, err := toml.DecodeFile(designFile, executable); err != nil {
		c.UI.Error(fmt.Sprintf(
			"Failed to decode design file %q: %s", designFile, err))
		return 1
	}

	errs := executable.Validate()
	if len(errs) != 0 {
		c.UI.Error(fmt.Sprintf(
			"%q is not valid template file. It has %d errors:", designFile, len(errs)))
		for _, err := range errs {
			c.UI.Error(fmt.Sprintf(
				"  * %s", err.Error()))
		}
		return ExitCodeFailed
	}

	c.UI.Info(fmt.Sprintf(
		"%q is valid template file.\n"+
			"You can generate cli project based on this file via `gcli apply` command", designFile))

	return ExitCodeOK

}

// Synopsis is a one-line, short synopsis of the command.
func (c *ValidateCommand) Synopsis() string {
	return "Validate design template file"
}

// Help is a long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (c *ValidateCommand) Help() string {
	helpText := `
Usage: gcli validate FILE

  Validate design template file which has required filed. If not it returns
  error and non zero value. 
`
	return strings.TrimSpace(helpText)
}
