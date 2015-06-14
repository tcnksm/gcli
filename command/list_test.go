package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestListCommand_implement(t *testing.T) {
	var _ cli.Command = &ListCommand{}
}
