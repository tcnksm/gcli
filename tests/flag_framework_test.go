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

func TestNew_flag_frameworks(t *testing.T) {

	tests := []struct {
		framework string
		expectOut string
	}{
		{
			framework: "flag",
			expectOut: "Usage of ",
		},
	}

	owner := "awesome_user_" + strconv.Itoa(int(time.Now().Unix()))
	cleanFunc, err := chdirSrcPath(owner)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanFunc()

	for _, tt := range tests {

		artifactBin := fmt.Sprintf("%s_grep", tt.framework)
		args := []string{
			"new",
			"-framework", tt.framework,
			"-owner", owner,
			"-flag=ignore:Bool:'Perform case insensitive matching'",
			"-flag=context:Int:'Print num lines of leading and trailing context'",
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
		cmd := exec.Command("./"+artifactBin, "-help")
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
