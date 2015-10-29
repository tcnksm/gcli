package skeleton

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string

	Framework  *Framework
	Executable *Executable

	// If WithTest is true, also generate test code.
	SkipTest bool

	// ArtifactCh is channel for info output
	ArtifactCh chan string

	// ErrCh is channel for error output
	ErrCh chan error

	// StaticDir contains static contents which will
	// be copied final output directory
	StaticDir string

	// Verbose enables logging output below INFO
	Verbose bool

	// LogWriter
	LogWriter io.Writer
}

// Generate generates code files from tempalte files.
func (s *Skeleton) Generate() <-chan struct{} {

	s.Debugf("Start generating")

	// doneCh is used to tell task it done to parent function
	doneCh := make(chan struct{})

	go func() {

		// Start generating base files
		doneBase := s.genBase()

		// Start generating command files
		doneTmpl := s.genTmpl()

		// Read static files and copy it on output directory.
		// Same name file which is generaed by gcli
		// will be overwrite by this
		doneStatic := s.copyStatic()

		<-doneStatic
		<-doneBase
		<-doneTmpl

		doneCh <- struct{}{}
	}()

	return doneCh
}

func (s *Skeleton) copyStatic() <-chan struct{} {
	s.Debugf("Start reading static files")
	doneCh := make(chan struct{})
	go func() {
		defer func() {
			doneCh <- struct{}{}
		}()

		// Ignore when staticDir is not provided.
		if s.StaticDir == "" {
			return
		}

		// Ignore when staticDir is not exist.
		if _, err := os.Stat(s.StaticDir); os.IsNotExist(err) {
			return
		}

		fis, err := ioutil.ReadDir(s.StaticDir)
		if err != nil {
			s.ErrCh <- err
			return
		}

		for _, fi := range fis {
			if fi.IsDir() {
				continue
			}

			srcFile := filepath.Join(s.StaticDir, fi.Name())
			src, err := os.Open(srcFile)
			if err != nil {
				s.ErrCh <- err
				return
			}
			defer src.Close()

			dstFile := filepath.Join(s.Path, fi.Name())

			// Create directory if necessary
			dir, _ := filepath.Split(dstFile)
			if dir != "" {
				if err := mkdir(dir); err != nil {
					s.ErrCh <- err
					return
				}
			}

			dst, err := os.Create(dstFile)
			if err != nil {
				s.ErrCh <- err
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, src); err != nil {
				s.Debugf(err.Error())
				s.ErrCh <- err
				return
			}

			s.ArtifactCh <- dstFile
		}
	}()
	return doneCh
}

func (s *Skeleton) genBase() <-chan struct{} {

	s.Debugf("Start generating base files")

	// doneCh is used to tell task it done
	doneCh := make(chan struct{})

	go func() {

		var wg sync.WaitGroup
		baseTmpls := CommonTemplates
		baseTmpls = append(baseTmpls, s.Framework.BaseTemplates...)
		for _, tmpl := range baseTmpls {
			s.Debugf("Use tempalte file: %s, output path tempalte string: %s",
				tmpl.Path, tmpl.OutputPathTmpl)

			if s.SkipTest && strings.HasSuffix(tmpl.Path, "_test.go.tmpl") {
				s.Debugf("Skip test tempalte file: %s", filepath.Base(tmpl.Path))
				continue
			}

			wg.Add(1)
			go func(tmpl Template) {
				defer wg.Done()
				tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
				outputPath, err := tmpl.Exec(s.Executable)
				if err != nil {
					s.ErrCh <- err
				}
				s.ArtifactCh <- outputPath
			}(tmpl)
		}

		// Wait until all task is done
		wg.Wait()

		// Tell doneCh about finishing generating
		doneCh <- struct{}{}
	}()

	return doneCh
}

func (s *Skeleton) genTmpl() <-chan struct{} {

	s.Debugf("Start generating command files")

	// doneCh is used to tell task it done
	doneCh := make(chan struct{})

	go func() {
		var wg sync.WaitGroup

		for _, cmd := range s.Executable.Commands {
			wg.Add(1)
			go func(cmd *Command) {
				defer wg.Done()
				for _, tmpl := range s.Framework.CommandTemplates {

					s.Debugf("Use tempalte file: %s, output path tempalte string: %s",
						tmpl.Path, tmpl.OutputPathTmpl)
					if s.SkipTest && strings.HasSuffix(tmpl.Path, "_test.go.tmpl") {
						s.Debugf("Skip test tempalte file: %s", tmpl.Path)
						continue
					}

					tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
					outputPath, err := tmpl.Exec(cmd)
					if err != nil {
						s.ErrCh <- err
					}
					s.ArtifactCh <- outputPath
				}
			}(cmd)
		}

		// Wait until all task is done.
		wg.Wait()

		doneCh <- struct{}{}
	}()

	return doneCh
}
