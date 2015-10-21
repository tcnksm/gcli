package skeleton

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyStatic(t *testing.T) {

	// Create temp static directory and file there
	staticDir, err := ioutil.TempDir("", "copy-static-src")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	fakeSrc, err := os.Create(filepath.Join(staticDir, "fake"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer fakeSrc.Close()

	content := "This is LICENSE file"
	if _, err = fakeSrc.WriteString(content); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Create temp output directory
	outputDir, err := ioutil.TempDir("", "copy-static-dst")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	artifactCh, errCh := make(chan string), make(chan error)
	fakeSkeleton := &Skeleton{
		Path:       outputDir,
		StaticDir:  staticDir,
		ErrCh:      errCh,
		ArtifactCh: artifactCh,
		LogWriter:  os.Stderr,
	}

	go func() {
		for _ = range artifactCh {
		}
	}()

	doneCh := fakeSkeleton.copyStatic()
	select {
	case <-doneCh:
	case err := <-fakeSkeleton.ErrCh:
		t.Fatalf("err: %s", err)
	}

	artifact, err := ioutil.ReadFile(filepath.Join(outputDir, "fake"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if string(artifact) != content {
		t.Fatalf("expects %q to be eq %q", string(artifact), content)
	}
}
