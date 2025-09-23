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
	stoFixtures = filepath.Join(filepath.Dir(b), FixturesDirectory, "sto")
)

func TestStore(t *testing.T) {
	err := filepath.WalkDir(stoFixtures,
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
			sto, err := OpenSTO(file)
			if err != nil {
				return err
			}
			if sto == nil {
				return fmt.Errorf("Parsed store is nil")
			}
			fixture, err := os.ReadFile(path + ".json")
			if err != nil {
				return err
			}
			expected := STO{}
			if err = json.Unmarshal(fixture, &expected); err != nil {
				return err
			}
			if !sto.Equal(&expected) {
				t.Fatalf("Result:\n%+v\n Does not match Expected:\n%+v\n", sto, expected)
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse store files: %+v", err)
	}
}
