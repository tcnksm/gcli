package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GitConfig(key string) string {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "config", "--global", "--get", "--null", key)
	cmd.Stdout = &stdout
	cmd.Stderr = ioutil.Discard

	if err := cmd.Run(); err != nil {
		return ""
	}

	return strings.TrimRight(stdout.String(), "\000")
}
