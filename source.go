package main

import (
	"os"
	"strings"
	"text/template"
)

type GoSource struct {
	Name     string
	Template template.Template
}

func (f GoSource) generate(appName string, definition Application) error {
	wr, err := os.Create(strings.Join([]string{appName, f.Name}, "/"))
	if err != nil {
		return err
	}

	defer wr.Close()
	return f.Template.Execute(wr, definition)
}
