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
	logFixtures = filepath.Join(filepath.Dir(b), FixturesDirectory, "log")
)

func TestLog(t *testing.T) {
	err := filepath.WalkDir(logFixtures,
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
			log, err := OpenLog(file)
			if err != nil {
				return err
			}
			if log == nil {
				return fmt.Errorf("Parsed Log is nil")
			}
			fixture, err := os.ReadFile(path + ".json")
			if err != nil {
				return err
			}
			expected := &Log{}
			if err = json.Unmarshal(fixture, expected); err != nil {
				return err
			}
			if !log.Equal(expected) {
				t.Fatalf("Result:\n%+v\n Does not match Expected:\n%+v\n", log, expected)
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse Log files, %+v", err)
	}
}
