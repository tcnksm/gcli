package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestApply(t *testing.T) {

	designFile, err := filepath.Abs("./fixtures/command-design.toml")
	if err != nil {
		t.Fatal(err)
	}

	// expect output when no flag/command provided
	expectOut := "Manage TODO tasks from command line (from design file)"

	// expect output when execute `add` command
	expectAddOutput := "Add new task"

	tests := []struct {
		framework string
		// is it include full description message
		// which is provided from outside
		fullDescription bool
	}{
		{
			framework:       "codegangsta_cli",
			fullDescription: true,
		},
		{
			framework:       "mitchellh_cli",
			fullDescription: false,
		},
		{
			framework:       "go_cmd",
			fullDescription: true,
		},
	}

	owner := "awesome_user_" + strconv.Itoa(int(time.Now().Unix()))
	cleanFunc, err := chdirSrcPath(owner)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanFunc()

	for _, tt := range tests {
		artifactBin := fmt.Sprintf("%s_todo_design_file", tt.framework)
		args := []string{
			"apply",
			"-framework", tt.framework,
			"-owner", owner,
			"-name", artifactBin,
			designFile,
		}

		output, err := runGcli(args)
		if err != nil {
			t.Fatal(err)
		}

		expect := "Successfully generated"
		if !strings.Contains(output, expect) {
			t.Fatalf("[%s] expect %q to contain %q", tt.framework, output, expect)
		}

		if err := os.Chdir(artifactBin); err != nil {
			t.Fatal(err)
		}

		if err := goGet(artifactBin); err != nil {
			t.Fatalf("[%s] Failed to run go get %s: %s", tt.framework, artifactBin, err)
		}

		if err := goVet(artifactBin); err != nil {
			t.Fatalf("[%s] Failed to run go vet %s: %s", tt.framework, artifactBin, err)
		}

		if err := goBuild(artifactBin); err != nil {
			t.Fatalf("[%s] Failed to run go build %s: %s", tt.framework, artifactBin, err)
		}

		binOutput := executeBin(artifactBin, []string{})
		if tt.fullDescription && !strings.Contains(binOutput, expectOut) {
			t.Errorf("[%s] expects %q to contain %q", tt.framework, binOutput, expectOut)
		}

		addOutput := executeBin(artifactBin, []string{"add"})
		if !strings.Contains(addOutput, expectAddOutput) {
			t.Errorf("[%s] expects %q to contain %q", tt.framework, addOutput, expectAddOutput)
		}

		// Back to src directory
		if err := os.Chdir(".."); err != nil {
			t.Fatal(err)
		}
	}
}
