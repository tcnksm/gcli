package skeleton

import (
	"io"
	"path/filepath"
	"text/template"
)

type Template struct {
	// Path is the path to this template
	Path string

	// contents is string contents of this template
	contents string
}

func NewTemplate(path string) (*Template, error) {
	template := &Template{Path: path}
	if err := template.init(); err != nil {
		return nil, err
	}

	return template, nil
}

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

// init read tempalte from path
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
