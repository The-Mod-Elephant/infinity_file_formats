package bg

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _  = runtime.Caller(0)
	basepath    = filepath.Dir(b)
	arefixtures = filepath.Join(basepath, "../fixtures", "are")
)

func TestArea(t *testing.T) {
	err := filepath.WalkDir(arefixtures,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			area, err := OpenArea(file)
			if err != nil {
				return err
			}
			if area == nil {
				return fmt.Errorf("Area nil failed")
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse area files, %+v", err)
	}
}
