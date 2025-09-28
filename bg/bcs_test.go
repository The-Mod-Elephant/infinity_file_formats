package bg

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

const BCS_FOLDER = "bcs"

var (
	bcsFixtures = filepath.Join(filepath.Dir(b), "../fixtures", BCS_FOLDER)
)

func TestBcs(t *testing.T) {
	err := filepath.WalkDir(bcsFixtures,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			println(d.Name())
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			bcs, err := OpenBcs(file)
			if err != nil {
				return err
			}
			if bcs == nil {
				return fmt.Errorf("Effects nil failed")
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse effects files, %+v", err)
	}
}
