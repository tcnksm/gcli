package skeleton

import (
	"path/filepath"
	"strings"
	"sync"
)

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string

	// Framework represent which cli package is used.
	// Framework ID is defined on framework_tempalte.go
	Framework int

	// If WithTest is true, also generate test code.
	SkipTest bool

	Executable *Executable
}

// Generate generates code files from tempalte files.
func (s *Skeleton) Generate() <-chan error {

	// Create error channel to return
	errCh := make(chan error)

	go func() {

		// Start generating base files
		errBaseCh := s.generateBaseFiles()

		// Start generating command files
		errCmndCh := s.generateCommandFiles()

		// Start generating custom files
		// which is generated from user defined templates
		errCstmCh := s.generateCustomFiles()

		// Merge all error channels until all channel is closed
		for err := range merge(errBaseCh, errCmndCh, errCstmCh) {
			errCh <- err
		}

		// Close channel after everything is Done.
		close(errCh)
	}()

	return errCh
}

func (s *Skeleton) generateBaseFiles() <-chan error {

	errCh := make(chan error)

	go func() {
		var wg sync.WaitGroup
		baseTmpls := CommonTemplates
		baseTmpls = append(baseTmpls, FrameworkTemplates(s.Framework)...)
		for _, tmpl := range baseTmpls {

			if s.SkipTest && strings.HasPrefix(tmpl.Path, "_test.go.tmpl") {
				continue
			}

			wg.Add(1)
			go func(tmpl Template) {
				defer wg.Done()
				tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
				if err := tmpl.Exec(s.Executable); err != nil {
					errCh <- err
				}
			}(tmpl)
		}

		// Wait until all task is done.
		wg.Wait()

		close(errCh)
	}()

	return errCh
}

func (s *Skeleton) generateCommandFiles() <-chan error {
	errCh := make(chan error)

	go func() {
		var wg sync.WaitGroup
		cmdTmpl, cmdTestTmpl := CommandTemplates(s.Framework)

		for _, cmd := range s.Executable.Commands {
			wg.Add(1)
			go func(tmpl Template, cmd Command) {
				defer wg.Done()
				tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
				if err := tmpl.Exec(cmd); err != nil {
					errCh <- err
				}
			}(cmdTmpl, cmd)

			if s.SkipTest {
				continue
			}

			wg.Add(1)
			go func(tmpl Template, cmd Command) {
				defer wg.Done()
				tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
				if err := tmpl.Exec(cmd); err != nil {
					errCh <- err
				}
			}(cmdTestTmpl, cmd)

		}

		// Wait until all task is done.
		wg.Wait()
		close(errCh)
	}()

	return errCh
}

func (s *Skeleton) generateCustomFiles() <-chan error {
	errCh := make(chan error)
	defer close(errCh)
	return errCh
}

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
