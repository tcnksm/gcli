package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

const EnvGcliPath = "GCLI_PATH"

// runGcliMux is mutex for running gcli because GOPATH
// is modified while test
var runGcliMux sync.Mutex

// gcli is executable path of gcli
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

// tmpGopath create temporary GOPATH.
// it returns its path and cleanFunction to remove directory.
// return error if any.
func tmpGopath() (string, func(), error) {
	gopath, err := ioutil.TempDir("./", "test-tmp-gopath-")
	if err != nil {
		return "", nil, err
	}

	absGopath, err := filepath.Abs(gopath)
	if err != nil {
		return "", nil, err
	}

	return absGopath, func() {
		if err := os.RemoveAll(absGopath); err != nil {
			panic(err)
		}
	}, nil
}

// setEnv set enviromental variables and return restore function.
func setEnv(key, val string) func() {

	preVal := os.Getenv(key)
	os.Setenv(key, val)

	return func() {
		os.Setenv(key, preVal)
	}
}

func createFiles(outDir string, files []string) error {
	// Check directory is exsit or not
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}

	if fi, _ := os.Stat(outDir); !fi.IsDir() {
		return fmt.Errorf("%q shoudl be directory", outDir)
	}

	for _, file := range files {
		newFile := filepath.Join(outDir, file)
		fi, err := os.Create(newFile)
		if err != nil {
			return err
		}
		defer fi.Close()
	}

	return nil
}

func checkFiles(outDir string, files []string) error {

	// Check directory is exsit or not
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}

	if fi, _ := os.Stat(outDir); !fi.IsDir() {
		return fmt.Errorf("%q shoudl be directory", outDir)
	}

	// Read dir files
	outputs, err := ioutil.ReadDir(outDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		isExist := false
		for _, fi := range outputs {
			if file == fi.Name() {
				isExist = true
				break
			}
		}

		if !isExist {
			return fmt.Errorf("%q is not exsit", file)
		}
	}

	return nil
}

func runGcli(runDir, gopath string, args []string) (string, error) {
	// Only one process can run gcli because it required
	// modifying GOPATH and changing directory.
	runGcliMux.Lock()
	defer runGcliMux.Unlock()

	// Modiry GOPATH while running gcli
	resetFunc := setEnv("GOPATH", gopath)
	defer resetFunc()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(gcli, args...)
	cmd.Dir = runDir
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

func runExecutable(bin string, args []string) string {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(bin, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	// cmd.Wait() returns error
	_ = cmd.Run()

	return stdout.String() + stderr.String()
}
