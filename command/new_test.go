package command

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestNewCommand_implement(t *testing.T) {
	var _ cli.Command = &NewCommand{}
}

func TestNewCommand(t *testing.T) {
	ui := new(cli.MockUi)
	c := &NewCommand{
		Meta: Meta{
			UI: ui,
		},
	}

	// Create temp directory to output file
	tmpDir, err := ioutil.TempDir("", "new-command")
	if err != nil {
		t.Fatal(err)
	}

	backFunc, err := TmpChdir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	defer backFunc()

	args := []string{"-F", "mitchellh_cli", "-owner", "deeeet", "todo"}
	if code := c.Run(args); code != 0 {
		t.Fatalf("bad status code: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	// TODO, inspect generated files
}

func TestNewCommand_directoryExist(t *testing.T) {
	ui := new(cli.MockUi)
	c := &NewCommand{
		Meta: Meta{
			UI: ui,
		},
	}

	// Create temp directory to output file
	tmpDir, err := ioutil.TempDir("", "new-command")
	if err != nil {
		t.Fatal(err)
	}

	backFunc, err := TmpChdir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	defer backFunc()

	// Create `todo` directory, same name as
	// application which is generated later step
	if err := os.Mkdir("todo", 0777); err != nil {
		t.Fatal(err)
	}

	args := []string{"todo"}
	if code := c.Run(args); code != 1 {
		t.Fatalf("bad status code: %d", code)
	}

	output := ui.ErrorWriter.String()
	expect := "Cannot create directory todo: file exists"
	if !strings.Contains(output, expect) {
		t.Errorf("expect %q to contain %q", output, expect)
	}
}
