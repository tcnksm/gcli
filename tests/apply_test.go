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

var testDesignFile = "./fixtures/command-design.toml"

var applyTests = []struct {
	framework string
}{
	{framework: "codegangsta_cli"},
	{framework: "mitchellh_cli"},
	{framework: "go_cmd"},
}

func TestApply(t *testing.T) {
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

	designFile, err := filepath.Abs(testDesignFile)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	for _, tt := range applyTests {
		name := fmt.Sprintf("%s_todo_design_file", tt.framework)
		args := []string{
			"apply",
			"-framework", tt.framework,
			"-owner", owner,
			"-name", name,
			designFile,
		}

		output, err := runGcli(baseDir, gopath, args)
		if err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		expectWarn := "WARNING: You are not in the directory gcli expects."
		if strings.Contains(output, expectWarn) {
			t.Fatalf("[%s] expects output not to contain %q", tt.framework, expectWarn)
		}

		expect := "Successfully generated"
		if !strings.Contains(output, expect) {
			t.Fatalf("[%s] expects output to contain %q", tt.framework, expect)
		}
	}
}

func TestApply_gotests(t *testing.T) {
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

	designFile, err := filepath.Abs(testDesignFile)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	for _, tt := range applyTests {
		name := fmt.Sprintf("%s_todo_design_file", tt.framework)
		args := []string{
			"apply",
			"-framework", tt.framework,
			"-owner", owner,
			"-name", name,
			designFile,
		}

		if _, err := runGcli(baseDir, gopath, args); err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		if err := goTests(filepath.Join(baseDir, name), gopath); err != nil {
			t.Fatalf("[%s] expects generated project to pass all go tests: \n\n %s", tt.framework, err)
		}
	}
}

func TestApply_vcs(t *testing.T) {
	t.Parallel()

	vcsHost := "bitbucket.org"
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

	designFile, err := filepath.Abs(testDesignFile)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	for _, tt := range applyTests {
		name := fmt.Sprintf("%s_todo_design_file", tt.framework)
		args := []string{
			"apply",
			"-framework", tt.framework,
			"-owner", owner,
			"-vcs", vcsHost,
			"-name", name,
			designFile,
		}

		output, err := runGcli(baseDir, gopath, args)
		if err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		expectWarn := "WARNING: You are not in the directory gcli expects."
		if strings.Contains(output, expectWarn) {
			t.Fatalf("expects output not to contain %q", expectWarn)
		}

		if err := goTests(filepath.Join(baseDir, name), gopath); err != nil {
			t.Fatalf("[%s] expects generated project to pass all go tests: \n\n %s", tt.framework, err)
		}
	}
}
