package helper

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"os"
)

// GoFmt runs `gofmt` to io.Reader and save it as file
// If something wrong it returns error.
func GoFmt(filename string, in io.Reader) error {

	if in == nil {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, filename, src, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.SortImports(fileSet, file)

	var buf bytes.Buffer
	tabWidth := 8
	printerMode := printer.UseSpaces | printer.TabIndent
	err = (&printer.Config{Mode: printerMode, Tabwidth: tabWidth}).Fprint(&buf, fileSet, file)
	if err != nil {
		return err
	}

	res := buf.Bytes()
	if !bytes.Equal(src, res) {
		err = ioutil.WriteFile(filename, res, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
