package bg

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var (
	stoFixtures = filepath.Join(filepath.Dir(b), "../fixtures", "sto")
)

func TestStore(t *testing.T) {
	err := filepath.WalkDir(stoFixtures,
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
			itm, err := OpenSTO(file)
			if err != nil {
				return err
			}
			if itm == nil {
				return fmt.Errorf("Parsed store is nil")
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse store files, %+v", err)
	}
}
