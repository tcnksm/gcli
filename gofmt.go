package main

import (
	"io/ioutil"
	"os/exec"
)

func GoFmt(dir string) error {
	cmd := exec.Command("gofmt", "-w", dir)
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = ioutil.Discard

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
