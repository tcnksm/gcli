package main

import (
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	args := []string{
		"--version",
	}

	gopath, cleanFunc, err := tmpGopath()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer cleanFunc()

	output, err := runGcli("./", gopath, args)
	if err != nil {
		t.Fatalf("expects %s to be nil", err)
	}

	// expect sample: gcli version v0.2.3
	expect := regexp.MustCompile(`gcli version v\d+\.\d+\.\d+`)
	if !expect.MatchString(output) {
		t.Fatalf("expects output not to contain %q %s", expect, output)
	}
}
