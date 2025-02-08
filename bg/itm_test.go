package bg

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var (
	itmFixtures = filepath.Join(filepath.Dir(b), "../fixtures", "itm")
)

func TestItem(t *testing.T) {
	err := filepath.WalkDir(itmFixtures,
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
			itm, err := OpenITM(file)
			if err != nil {
				return err
			}
			if itm == nil {
				return fmt.Errorf("Effects nil failed")
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse effects files, %+v", err)
	}
}
