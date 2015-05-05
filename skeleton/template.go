package skeleton

import (
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

type Template struct {
	// Path is the path to this template
	Path string

	// Output is the output name of this template
	Output string

	// contents is string contents of this template
	contents string
}

func NewTemplate(path string) (*Template, error) {
	template := &Template{
		Path:   path,
		Output: strings.TrimRight(filepath.Base(path), ".tmpl"),
	}

	if err := template.init(); err != nil {
		return nil, err
	}

	return template, nil
}

// Must is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil. It is intended for use in variable initializations
func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

// Execute applies a parsed template to the specified data object
func (t *Template) Execute(wr io.Writer, data *Executable) error {
	name := filepath.Base(t.Path)
	funcs := funcMap()

	tmpl, err := template.New(name).Funcs(funcs).Parse(t.contents)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(wr, data); err != nil {
		return err
	}

	return nil
}

// init read tempalte from path and set Output name by its basename
func (t *Template) init() error {
	contents, err := Asset(t.Path)
	if err != nil {
		return err
	}

	t.contents = string(contents)
	return nil
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"date": dateFunc(),
	}
}
