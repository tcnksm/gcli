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

	"github.com/tcnksm/gcli/skeleton"
)

var commandTests = []struct {
	framework  string
	args       []string
	expectHelp string
}{
	{
		framework:  "codegangsta_cli",
		args:       []string{},
		expectHelp: "[global options] command [command options] [arguments...]",
	},
	{
		framework:  "mitchellh_cli",
		args:       []string{},
		expectHelp: "[--version] [--help] <command> [<args>]",
	},
	{
		framework:  "go_cmd",
		args:       []string{},
		expectHelp: "help [command]\" for more information about a command.",
	},
}

func TestNew_command(t *testing.T) {
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

	for _, tt := range commandTests {

		name := fmt.Sprintf("%s_todo", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-flag=verbose:bool:'Run verbose mode'",
			"-flag=username:string:'Username'",
			"-flag=dry-run:string:'Dry-run mode'",
			"-command=add:'Add new task'",
			"-command=list:'List tasks'",
			"-command=change-state:'Change task state'",
			"-command=delete:'Delete specified task'",
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

func TestNew_command_unidealPath(t *testing.T) {
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

	for _, tt := range commandTests {

		name := fmt.Sprintf("%s_todo", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-command=add:'Add new task'",
			"-command=list:'List tasks'",
			"-command=delete:'Delete specified task'",
			name,
		}

		output, err := runGcli("./", gopath, args)
		if err != nil {
			t.Fatalf("[%s] expects %s to be nil", tt.framework, err)
		}

		expectWarn := "WARNING: You are not in the directory gcli expects."
		if !strings.Contains(output, expectWarn) {
			t.Fatalf("[%s] expects output to contain %q", tt.framework, expectWarn)
		}

		expect := "Successfully generated"
		if !strings.Contains(output, expect) {
			t.Fatalf("[%s] expects output to contain %q", tt.framework, expect)
		}
	}
}

// TestNew_command_gocommand tests that the generated project
// is go-buildable and it passes the go-test and go-vet.
func TestNew_command_gotests(t *testing.T) {
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

	for _, tt := range commandTests {

		name := fmt.Sprintf("%s_todo", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-command=add:'Add new task'",
			"-command=list:'List tasks'",
			"-command=change-state:'Change task state'",
			"-command=delete:'Delete specified task'",
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

func TestNew_command_checkOutputs(t *testing.T) {
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

	staticFiles := []string{"StaticA", "StaticB", "StaticC"}
	staticDir, err := ioutil.TempDir("", "gcli-test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if err := createFiles(staticDir, staticFiles); err != nil {
		t.Fatalf("err: %s", err)
	}

	for _, tt := range commandTests {

		name := fmt.Sprintf("%s_todo", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-static-dir", staticDir,
			"-command=add:'Add new task'",
			"-command=list:'List tasks'",
			"-command=change-state:'Change task state'",
			"-command=delete:'Delete specified task'",
			name,
		}

		if _, err := runGcli(baseDir, gopath, args); err != nil {
			t.Fatalf("err: %s", err)
		}

		targets := staticFiles

		// Collecting common files
		for _, tmpl := range skeleton.CommonTemplates {
			// NOTE: OutputPathTmpl of common template is same as final output name
			// and not changed by templating
			targets = append(targets, tmpl.OutputPathTmpl)
		}

		if err := checkFiles(filepath.Join(baseDir, name), targets); err != nil {
			t.Fatal(err)
		}
	}
}

func TestNew_command_vcs(t *testing.T) {
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

	for _, tt := range commandTests {

		name := fmt.Sprintf("%s_todo", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-vcs", vcsHost,
			"-command=add:'Add new task'",
			"-command=list:'List tasks'",
			"-command=change-state:'Change task state'",
			"-command=delete:'Delete specified task'",
			name,
		}

		output, err := runGcli(baseDir, gopath, args)
		if err != nil {
			t.Fatalf("err: %s", err)
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
