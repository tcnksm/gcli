package skeleton

import (
	"os"
	"path/filepath"
)

type Skeleton struct {
	// Path is where skeleton is generated
	Path string

	// Executable
	Executable *Executable

	// Framework
	Framework int
}

// Create generate codes. Create directory and generate code with tempalte file
func (s *Skeleton) Generate() error {
	if err := s.prepareDir(); err != nil {
		return err
	}

	path := "resource/tmpl/common/CHANGELOG.md.tmpl"
	tempalte, err := NewTemplate(path)
	if err != nil {
		return err
	}

	path = filepath.Join(s.Path, "CHANGELOG.md")
	wr, _ := os.Create(path)
	defer wr.Close()

	if err := tempalte.Execute(wr, s.Executable); err != nil {
		return err
	}

	return nil
}

// PrepareDir creates default directories.
func (s *Skeleton) prepareDir() error {
	path := filepath.Join(s.Path, "command")
	// If using flag pacakge, it doesn't require command directory.
	if s.Framework == Framework_flag {
		path = filepath.Join(s.Path)
	}

	return os.MkdirAll(path, 0766)
}
