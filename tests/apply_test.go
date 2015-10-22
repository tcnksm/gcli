package main

import (
	"fmt"
	"io/ioutil"
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

	staticFiles := []string{"StaticA", "StaticB"}
	staticDir, err := ioutil.TempDir("", "gcli-test")
	if err != nil {
		t.Fatal(err)
	}
	createFiles(staticDir, staticFiles)

	for _, tt := range tests {
		artifactBin := fmt.Sprintf("%s_todo_design_file", tt.framework)
		args := []string{
			"apply",
			"-framework", tt.framework,
			"-owner", owner,
			"-name", artifactBin,
			"-static-dir", staticDir,
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

		if err := checkFiles(artifactBin, staticFiles); err != nil {
			t.Fatal(err)
		}

		if err := goTests(artifactBin); err != nil {
			t.Fatal(err)
		}

		if err := os.Chdir(artifactBin); err != nil {
			t.Fatal(err)
		}

		binOutput := executeBin(artifactBin, []string{})
		if tt.fullDescription && !strings.Contains(binOutput, expectOut) {
			t.Errorf("[%s] expects %q to contain %q", tt.framework, binOutput, expectOut)
		}

		// Need to fix after https://github.com/BurntSushi/toml/pull/90 is fixed
		// addOutput := executeBin(artifactBin, []string{"add"})
		// if !strings.Contains(addOutput, expectAddOutput) {
		// 	t.Errorf("[%s] expects %q to contain %q", tt.framework, addOutput, expectAddOutput)
		// }
		_ = expectAddOutput

		// Back to src directory
		if err := os.Chdir(".."); err != nil {
			t.Fatal(err)
		}
	}
}
