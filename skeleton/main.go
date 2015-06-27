package skeleton

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"sync"
)

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string

	// If WithTest is true, also generate test code.
	SkipTest bool

	Framework  *Framework
	Executable *Executable

	// OutCh is channel for info output
	OutCh chan string

	// ErrCh is channel for error output
	ErrCh chan error

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
		doneBaseCh := s.generateBaseFiles()

		// Start generating command files
		doneCmdCh := s.generateCommandFiles()

		// Start generating custom files
		// which is generated from user defined templates
		// s.generateCustomFiles()

		<-doneBaseCh
		<-doneCmdCh

		doneCh <- struct{}{}
	}()

	return doneCh
}

func (s *Skeleton) generateBaseFiles() <-chan struct{} {

	s.Debugf("Start generating base files")

	// doneCh is used to tell task it done
	doneCh := make(chan struct{})

	go func() {

		var wg sync.WaitGroup
		baseTmpls := CommonTemplates
		baseTmpls = append(baseTmpls, s.Framework.BaseTemplates...)
		for _, tmpl := range baseTmpls {

			s.Debugf("Use tempalte file: %s, output path tempalte string: %s", tmpl.Path, tmpl.OutputPathTmpl)
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
				s.OutCh <- fmt.Sprintf("Create %s", outputPath)
			}(tmpl)
		}

		// Wait until all task is done
		wg.Wait()

		// Tell doneCh about finishing generating
		doneCh <- struct{}{}
	}()

	return doneCh
}

func (s *Skeleton) generateCommandFiles() <-chan struct{} {

	s.Debugf("Start generating command files")

	// doneCh is used to tell task it done
	doneCh := make(chan struct{})

	go func() {
		var wg sync.WaitGroup

		for _, cmd := range s.Executable.Commands {
			wg.Add(1)
			go func(cmd Command) {
				defer wg.Done()
				for _, tmpl := range s.Framework.CommandTemplates {

					s.Debugf("Use tempalte file: %s, output path tempalte string: %s", tmpl.Path, tmpl.OutputPathTmpl)
					if s.SkipTest && strings.HasSuffix(tmpl.Path, "_test.go.tmpl") {
						s.Debugf("Skip test tempalte file: %s", tmpl.Path)
						continue
					}

					tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
					outputPath, err := tmpl.Exec(cmd)
					if err != nil {
						s.ErrCh <- err
					}
					s.OutCh <- fmt.Sprintf("Create %s", outputPath)
				}
			}(cmd)
		}

		// Wait until all task is done.
		wg.Wait()

		doneCh <- struct{}{}
	}()

	return doneCh
}

// func (s *Skeleton) generateCustomFiles() <-chan struct{} {
// 	doneCh := make(chan struct{})
// 	defer func() {
// 		doneCh <- struct{}{}
// 	}()
// 	return doneCh
// }

// merge merges error channels and sends them to union channel
func merge(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	out := make(chan error)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(errCh <-chan error) {
			defer wg.Done()
			for err := range errCh {
				out <- err
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
