package main

import (
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestIsGoFile(t *testing.T) {
	RegisterTestingT(t)

	f, _ := os.Stat("cli-init.go")
	Expect(isGofile(f)).To(BeTrue())

	f, _ = os.Stat("README.md")
	Expect(isGofile(f)).To(BeFalse())
}
