package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNew_command_frameworks(t *testing.T) {

	tests := []struct {
		framework string
		expectOut string
	}{
		{
			framework: "codegangsta_cli",
			expectOut: "[global options] command [command options] [arguments...]",
		},
		{
			framework: "mitchellh_cli",
			expectOut: "[--version] [--help] <command> [<args>]",
		},
		{
			framework: "go_cmd",
			expectOut: "help [command]\" for more information about a command.",
		},
	}

	owner := "awesome_user_" + strconv.Itoa(int(time.Now().Unix()))
	cleanFunc, err := chdirSrcPath(owner)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanFunc()

	for _, tt := range tests {

		artifactBin := fmt.Sprintf("%s_todo", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-flag=verbose:bool:'Run verbose mode'",
			"-flag=username:string:'Username'",
			"-command=add:'Add new task'",
			"-command=list:'List tasks'",
			"-command=delete:'Delete specified task'",
			artifactBin,
		}

		output, err := runGcli(args)
		if err != nil {
			t.Fatal(err)
		}

		expect := "Successfully generated"
		if !strings.Contains(output, expect) {
			t.Fatalf("[%s] expect %q to contain %q", tt.framework, output, expect)
		}

		if err := goTests(artifactBin); err != nil {
			t.Fatal(err)
		}

		if err := os.Chdir(artifactBin); err != nil {
			t.Fatal(err)
		}

		var stdout, stderr bytes.Buffer
		cmd := exec.Command("./" + artifactBin)
		cmd.Stderr = &stderr
		cmd.Stdout = &stdout

		// cmd.Wait() returns error
		_ = cmd.Run()

		output = stdout.String() + stderr.String()
		// t.Logf("%s \n\n%s", tt.framework, output)
		if !strings.Contains(output, tt.expectOut) {
			t.Errorf("[%s] expects %q to contain %q", tt.framework, output, tt.expectOut)
		}

		// Back to src directory
		if err := os.Chdir(".."); err != nil {
			t.Fatal(err)
		}
	}
}
