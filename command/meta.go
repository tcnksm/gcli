package command

import (
	"bufio"
	"flag"
	"io"

	"github.com/mitchellh/cli"
)

// ExitCodes
const (
	ExitCodeOK     int = 0
	ExitCodeFailed int = 1
)

const (
	// defaultFrameworkString is default cli framework name
	defaultFrameworkString = "codegangsta_cli"
)

// Meta contain the meta-option that nealy all subcommand inherited.
type Meta struct {
	UI cli.Ui
}

// NewFlagSet generates commom flag.FlagSet
func (m *Meta) NewFlagSet(name string, helpText string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)

	// Set usage function
	flags.Usage = func() { m.UI.Error(helpText) }

	// Set error output to Meta.UI.Error
	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	flags.SetOutput(errW)

	go func() {
		for errScanner.Scan() {
			m.UI.Error(errScanner.Text())
		}
	}()

	return flags
}
