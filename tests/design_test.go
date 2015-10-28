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

var designTests = []struct {
	framework string
}{
	{framework: "codegangsta_cli"},
	{framework: "mitchellh_cli"},
	{framework: "go_cmd"},
}

func TestDesign(t *testing.T) {
	t.Parallel()

	vcsHost := "github.com"
	owner := "awesome_user_" + strconv.Itoa(int(time.Now().Unix()))

	gopath, cleanFunc, err := tmpGopath()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer cleanFunc()

	baseDir := filepath.Join(gopath, "src", vcsHost, owner)
	if err := os.MkdirAll(baseDir, 0777); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Create design File
	for _, tt := range designTests {
		name := fmt.Sprintf("%s-git", tt.framework)
		designFile := fmt.Sprintf("%s-design-test.toml", name)
		designArgs := []string{
			"design",
			"-owner", owner,
			"-framework", tt.framework,
			"-command=add:'Add file contents to the index'",
			"-command=commit:'Record changes to the repository'",
			"-command=push:'Update remote refs along with associated objects'",
			"-command=pull-request:'Open a pull request on GitHub'",
			"-output", designFile,
			name,
		}

		output, err := runGcli(baseDir, gopath, designArgs)
		if err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		expect := "====> Successfully generated"
		if !strings.Contains(output, expect) {
			t.Fatalf("[%s] expects output to contain %q", tt.framework, expect)
		}

		// Check design file is exist or not
		if _, err := os.Stat(filepath.Join(baseDir, designFile)); os.IsNotExist(err) {
			t.Fatalf("[%s] expects %q to be exist", tt.framework, designFile)
		}
	}
}

func TestDesign_gotests(t *testing.T) {
	t.Parallel()

	vcsHost := "github.com"
	owner := "awesome_user_" + strconv.Itoa(int(time.Now().Unix()))

	gopath, cleanFunc, err := tmpGopath()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer cleanFunc()

	baseDir := filepath.Join(gopath, "src", vcsHost, owner)
	if err := os.MkdirAll(baseDir, 0777); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Create design File
	for _, tt := range designTests {
		name := fmt.Sprintf("%s-git", tt.framework)
		designFile := fmt.Sprintf("%s-design-test.toml", name)
		designArgs := []string{
			"design",
			"-owner", owner,
			"-framework", tt.framework,
			"-command=add:'Add file contents to the index'",
			"-command=commit:'Record changes to the repository'",
			"-command=push:'Update remote refs along with associated objects'",
			"-command=pull-request:'Open a pull request on GitHub'",
			"-output", designFile,
			name,
		}

		if _, err := runGcli(baseDir, gopath, designArgs); err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		applyArgs := []string{
			"apply",
			"-name", name,
			filepath.Join(baseDir, designFile),
		}

		if _, err := runGcli(baseDir, gopath, applyArgs); err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		if err := goTests(filepath.Join(baseDir, name), gopath); err != nil {
			t.Fatalf("[%s] expects generated project to pass all go tests: \n\n %s", tt.framework, err)
		}
	}
}
