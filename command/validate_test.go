package command

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/mitchellh/cli"
)

func TestValidateCommand_implement(t *testing.T) {
	var _ cli.Command = &ValidateCommand{}
}

func TestValidateCommand(t *testing.T) {

	fixture := "./fixtures"

	tests := []struct {
		input    string
		exitCode int
	}{
		{
			input:    "valid-design.toml",
			exitCode: ExitCodeOK,
		},

		{
			input:    "invalid-design.toml",
			exitCode: ExitCodeFailed,
		},
	}

	for i, tt := range tests {
		ui := new(cli.MockUi)
		c := &ValidateCommand{
			Meta: Meta{
				UI: ui,
			},
		}
		input := filepath.Join(fixture, tt.input)
		if code := c.Run([]string{input}); code != tt.exitCode {
			t.Fatalf("#%d bad status code: %d, expects %d\n\n%s", i, code, tt.exitCode, ui.ErrorWriter.String())
		}
		fmt.Println(ui.ErrorWriter.String())
	}
}
