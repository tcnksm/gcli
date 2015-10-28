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

var flagTests = []struct {
	framework  string
	args       []string
	expectHelp string
}{
	{
		framework:  "flag",
		args:       []string{"-h"},
		expectHelp: "Usage of ",
	},
}

func TestNew_flag(t *testing.T) {
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

	for _, tt := range flagTests {

		name := fmt.Sprintf("%s_grep", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-flag=ignore-case:Bool:'Perform case insensitive matching'",
			"-flag=context:Int:'Print num lines of leading and trailing context'",
			name,
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

// TestNew_flag_goflag tests that the generated project
// is go-buildable and it passes the go-test and go-vet.
func TestNew_flag_gotests(t *testing.T) {
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

	for _, tt := range flagTests {

		name := fmt.Sprintf("%s_grep", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-flag=ignore-case:Bool:'Perform case insensitive matching'",
			"-flag=context:Int:'Print num lines of leading and trailing context'",
			name,
		}

		if _, err := runGcli(baseDir, gopath, args); err != nil {
			t.Fatalf("err: %s", err)
		}

		if err := goTests(filepath.Join(baseDir, name), gopath); err != nil {
			t.Fatalf("[%s] expects generated project to pass all go tests: \n\n %s", tt.framework, err)
		}

		// Also run executable and check its output.
		// This test should be seaprated from this test.
		// But it has costs to run go-get multiple times.
		output := runExecutable(filepath.Join(baseDir, name, name), tt.args)
		if !strings.Contains(output, tt.expectHelp) {
			t.Errorf("[%s] expects %q to contain %q", tt.framework, output, tt.expectHelp)
		}
	}
}
