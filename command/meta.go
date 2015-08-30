package command

import "github.com/mitchellh/cli"

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
