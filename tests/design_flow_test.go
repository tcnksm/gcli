package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDesignFlow(t *testing.T) {
	// Test generating design file, validate it and generate
	// cli project from it (Testing all work flow).
	// let's create git interface.

	artifactBin := "mygit"

	owner := "awesome_user_" + strconv.Itoa(int(time.Now().Unix()))
	cleanFunc, err := chdirSrcPath(owner)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanFunc()

	// Create design File
	designFile := fmt.Sprintf("%s-design-test.toml", artifactBin)
	designArgs := []string{
		"design",
		"-owner", owner,
		"-framework", "mitchellh_cli",
		"-command=add:'Add file contents to the index'",
		"-command=commit:'Record changes to the repository'",
		"-command=push:'Update remote refs along with associated objects'",
		"-output", designFile,
		artifactBin,
	}

	if _, err := runGcli(designArgs); err != nil {
		t.Fatal(err)
	}

	// Check design file is exist or not
	if _, err := os.Stat(designFile); os.IsNotExist(err) {
		t.Fatal(err)
	}

	// Validate design File
	validateArgs := []string{
		"validate",
		designFile,
	}

	if _, err := runGcli(validateArgs); err != nil {
		t.Fatal(err)
	}

	// Apply to genearte cli project
	applyArgs := []string{
		"apply",
		designFile,
	}

	output, err := runGcli(applyArgs)
	if err != nil {
		t.Fatal(err)
	}

	expect := "Successfully generated"
	if !strings.Contains(output, expect) {
		t.Fatalf("Expect %q to contain %q", output, expect)
	}

	if err := goTests(artifactBin); err != nil {
		t.Fatalf("Failed to run go tests in %s: %s", artifactBin, err)
	}
}
