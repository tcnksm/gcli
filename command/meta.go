package command

import (
	"bufio"
	"flag"
	"io"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/go-homedir"
)

// ExitCodes
const (
	ExitCodeOK     int = 0
	ExitCodeFailed int = 1
)

const (
	// EnvGoPath is env name of GOPATH
	EnvGoPath = "GOPATH"
)

const (
	// DefaultVCSHost is the default VCS host name.
	DefaultVCSHost = "github.com"

	// DefaultLocalDir is the default path to store directory.
	DefaultLocalDir       = "~/.gcli.d"
	DefaultLocalStaticDir = "static"

	// defaultFrameworkString is default cli framework name
	defaultFrameworkString = "codegangsta_cli"
)

// Meta contain the meta-option that nealy all subcommand inherited.
type Meta struct {
	UI cli.Ui
}

// LocalDir returns the local directory for storing user defined data.
func (m *Meta) LocalDir() (string, error) {
	return homedir.Expand(DefaultLocalDir)
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
