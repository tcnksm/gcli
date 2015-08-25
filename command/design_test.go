package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/cli"
	"github.com/tcnksm/gcli/skeleton"
)

func TestDesignCommand_implement(t *testing.T) {
	var _ cli.Command = &DesignCommand{}
}

func TestDesignCommand(t *testing.T) {
	ui := new(cli.MockUi)
	c := &DesignCommand{
		Meta: Meta{
			UI: ui,
		},
	}

	// Create temp directory to output file
	tmpDir, err := ioutil.TempDir("", "design_command")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	name := "todo"
	if code := c.Run([]string{name}); code != 0 {
		t.Fatalf("bad status code: %d\n\n%s", code, ui.ErrorWriter.String())
	}

	// Inspect generated file
	outputFile := filepath.Join(tmpDir, fmt.Sprintf(defaultOutputFmt, name))
	executable := skeleton.NewExecutable()

	if _, err := toml.DecodeFile(outputFile, executable); err != nil {
		t.Fatal(err)
	}

	if executable.Name != name {
		t.Errorf("expects %q to be eq %q", executable.Name, name)
	}

	if executable.Version != skeleton.DefaultVersion {
		t.Errorf("expects %q to be eq %q", executable.Version, skeleton.DefaultVersion)
	}

	if executable.FrameworkStr != defaultFrameworkString {
		t.Errorf("expects %q to be eq %q", executable.FrameworkStr, defaultFrameworkString)
	}
}

func TestDesignCommand_fileExist(t *testing.T) {
	ui := new(cli.MockUi)
	c := &DesignCommand{
		Meta: Meta{
			UI: ui,
		},
	}

	// Create temp directory to output file
	tmpDir, err := ioutil.TempDir("", "apply-command")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	name := fmt.Sprintf(defaultOutputFmt, "todo")
	if _, err := os.Create(name); err != nil {
		t.Fatal(err)
	}

	if code := c.Run([]string{"todo"}); code != 1 {
		t.Fatalf("bad status code: %d", code)
	}

	output := ui.ErrorWriter.String()
	expect := fmt.Sprintf("Cannot create design file %s: file exists", name)
	if !strings.Contains(output, expect) {
		t.Errorf("expect %q to contain %q", output, expect)
	}
}
