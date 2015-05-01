package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestNewCommand_implement(t *testing.T) {
	var _ cli.Command = &NewCommand{}
}
