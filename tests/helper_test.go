package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const EnvGcliPath = "GCLI_PATH"

// gcli is executable path
var gcli string

func init() {
	gcli = os.Getenv(EnvGcliPath)
	if gcli == "" {
		gcli = "../bin/gcli"
	}

	// Should be absolute path so that we can change dir
	var err error
	gcli, err = filepath.Abs(gcli)
	if err != nil {
		panic(err)
	}
}

// chdirSrcPath changes dirctory to $GOPATH/src/github.com/owner/
// It returns cleanup script to delete directory
func chdirSrcPath(owner string) (func(), error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return nil, fmt.Errorf("can't found GOPATH env var")
	}

	srcPath := filepath.Join(gopath, "src", "github.com", owner)
	if _, err := os.Stat(srcPath); os.IsExist(err) {
		// TODO
		panic(err)
	}

	if err := os.MkdirAll(srcPath, 0777); err != nil {
		return nil, err
	}

	if err := os.Chdir(srcPath); err != nil {
		return nil, err
	}

	return func() {
		if err := os.RemoveAll(srcPath); err != nil {
			panic(err)
		}
	}, nil
}

// executeBin execute command and return output
func executeBin(bin string, args []string) string {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("./"+bin, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	// cmd.Wait() returns error
	_ = cmd.Run()
	return stdout.String() + stderr.String()
}

// runGcli runs gcli and return its stdout. If failed, returns error.
func runGcli(args []string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(gcli, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("failed to execute: %s\n\n %s", err, stderr.String())
	}

	return stdout.String(), nil

}

// goGet runs go get on current directory. If failed, returns error.
func goGet(output string) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "get", "-v", "-d", "-t", "./...")
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute: %s\n\n %s", err, stderr.String())
	}

	return nil
}

// goBuild runs go build on current directory. If failed, returns error.
func goBuild(output string) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "build", "-o", output)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute: %s\n\n %s", err, stderr.String())
	}

	return nil
}

// goVet runs go vet on current directory. If failed, returns error.
func goVet(output string) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "vet", "./...")
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute: %s\n\n %s", err, stderr.String())
	}

	return nil
}
