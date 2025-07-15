package bg

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var (
	effFixtures = filepath.Join(filepath.Dir(b), "../fixtures", "eff")
)

func TestEffect(t *testing.T) {
	err := filepath.WalkDir(effFixtures,
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
			effv1, effv2, err := OpenEff(file)
			if err != nil {
				return err
			}
			if effv1 == nil && effv2 == nil {
				return fmt.Errorf("Parsed effect is nil")
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse effects files, %+v", err)
	}
}
