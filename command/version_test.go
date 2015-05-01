package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestVersionCommand_implement(t *testing.T) {
	var _ cli.Command = &VersionCommand{}
}
