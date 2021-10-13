package genhelper

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/joncalhoun/pipe"
)

func TemplatePipe(t *template.Template, data interface{}) (io.ReadCloser, error) {
	rc, wc, errCh := pipe.Commands(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)
	go func() {
		select {
		case err, ok := <-errCh:
			if ok && err != nil {
				log.Fatalf("error in pipe: %v", err)
			}
		}
	}()

	err := t.Execute(wc, data)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	return rc, wc.Close()
}

func WriteToFile(src io.Reader, out string) (int64, error) {
	dst, err := os.Create(out)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer dst.Close()

	writtenBytes, err := io.Copy(dst, src)
	if err != nil {
		return 0, fmt.Errorf("error writing file: %v", err)
	}

	return writtenBytes, dst.Sync()
}
