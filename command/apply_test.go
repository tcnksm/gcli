package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestApplyCommand_implement(t *testing.T) {
	var _ cli.Command = &ApplyCommand{}
}
