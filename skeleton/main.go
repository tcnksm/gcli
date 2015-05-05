package skeleton

import (
	"os"
	"path/filepath"
	"sync"
)

type Skeleton struct {
	// Path is where skeleton is generated
	Path string

	// Executable
	Executable *Executable
}

// Create generate codes. Create directory and generate code with tempalte file
func (s *Skeleton) Generate() error {

	if err := s.prepareDir(); err != nil {
		return err
	}

	tmplList := CommonTemplates()
	tmplList = append(tmplList, FrameworkTemplates(s.Executable.Framework)...)
	doneCh, errCh := s.processTemplates(tmplList)
LOOP:
	for {
		select {
		case <-doneCh:
			break LOOP
		case err := <-errCh:
			// if at least one error is happend, terminate tempating
			return err
		}
	}
	return nil
}

// processTemplates run template execution of all templates in parallel.
func (s *Skeleton) processTemplates(tmpls []string) (<-chan bool, <-chan error) {

	// Channels to tell process state.
	doneCh, errCh := make(chan bool), make(chan error)

	// Block until all template execution is done.
	var wg sync.WaitGroup

	go func() {
		for _, path := range tmpls {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()

				tmpl := Must(NewTemplate(path))

				tempalte, err := NewTemplate(tmpl.Path)
				if err != nil {
					errCh <- err
					return
				}

				outPath := filepath.Join(s.Path, tmpl.Output)
				wr, err := os.Create(outPath)
				if err != nil {
					errCh <- err
					return
				}
				defer wr.Close()

				if err := tempalte.Execute(wr, s.Executable); err != nil {
					errCh <- err
					return
				}
			}(path)
		}

		wg.Wait()
		doneCh <- true
	}()

	return doneCh, errCh
}

// PrepareDir creates default directories.
func (s *Skeleton) prepareDir() error {
	path := filepath.Join(s.Path, "command")

	if s.Executable.Framework == Framework_flag {
		path = filepath.Join(s.Path)
	}

	if s.Executable.Framework == Framework_tcnksm_flag {
		path = filepath.Join(s.Path)
	}

	return os.MkdirAll(path, 0766)
}
