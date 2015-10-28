package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var goCmdMux sync.Mutex

// goTests runs go-get, go-build, go-test and goVet command in
// provided directory. If any, returns error.
func goTests(runDir, gopath string) error {

	goFuncs := []func(runDir, gopath string) error{
		goGet,
		goBuild,
		goTest,
		goVet,
	}

	for _, gf := range goFuncs {
		err := gf(runDir, gopath)
		if err != nil {
			return err
		}
	}

	return nil
}

// goGet runs go get on current directory. If failed, returns error.
func goGet(runDir, gopath string) error {
	goCmdMux.Lock()
	defer goCmdMux.Unlock()

	// Check directory is exsit or not
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}

	// Modiry GOPATH while running go command
	resetFunc := setEnv("GOPATH", gopath)
	defer resetFunc()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "get", "-v", "-d", "-t", "./...")
	cmd.Dir = runDir
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start `go get`: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute `go get`: %s\n\n %s", err, stderr.String())
	}

	return nil
}

// goBuild runs go build on current directory. If failed, returns error.
func goBuild(runDir, gopath string) error {
	goCmdMux.Lock()
	defer goCmdMux.Unlock()

	// Check directory is exsit or not
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}

	// Modiry GOPATH while running go command
	resetFunc := setEnv("GOPATH", gopath)
	defer resetFunc()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "build", "-o", filepath.Base(runDir))
	cmd.Dir = runDir
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start `go build`: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute `go build`: %s\n\n %s", err, stderr.String())
	}

	return nil
}

// goTest runs go test on current directory. If failed, returns error.
func goTest(runDir, gopath string) error {
	goCmdMux.Lock()
	defer goCmdMux.Unlock()

	// Check directory is exsit or not
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}

	// Modiry GOPATH while running go command
	resetFunc := setEnv("GOPATH", gopath)
	defer resetFunc()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = runDir
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start `go test`: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute `go test`: %s\n\n %s", err, stderr.String())
	}

	return nil
}

// goVet runs go vet on current directory. If failed, returns error.
func goVet(runDir, gopath string) error {
	goCmdMux.Lock()
	defer goCmdMux.Unlock()

	// Check directory is exsit or not
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}

	// Modiry GOPATH while running go command
	resetFunc := setEnv("GOPATH", gopath)
	defer resetFunc()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = runDir
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start `go vet`: %s\n\n %s", err, stderr.String())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to execute `go vet`: %s\n\n %s", err, stderr.String())
	}

	return nil
}
