package bg

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var (
	creFixtures = filepath.Join(filepath.Dir(b), "../fixtures", "cre")
)

func TestCreatures(t *testing.T) {
	err := filepath.WalkDir(creFixtures,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() || filepath.Ext(d.Name()) == ".json" {
				return nil
			}
			if err != nil {
				return err
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			cre, err := OpenCre(file)
			if err != nil {
				return err
			}
			if cre == nil {
				return fmt.Errorf("Parsed store is nil")
			}
			fixture, err := os.ReadFile(path + ".json")
			if err != nil {
				return err
			}
			expected := &CRE{}
			if err = json.Unmarshal(fixture, expected); err != nil {
				return err
			}
			if !cre.Equal(expected) {
				t.Fatalf("Result:\n%+v\n Does not match Expected:\n%+v\n", cre, expected)
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse effects files, %+v", err)
	}
}
