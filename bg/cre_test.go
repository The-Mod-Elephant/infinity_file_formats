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
	creFixtures = filepath.Join(filepath.Dir(b), FixturesDirectory, "cre")
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

func TestCreaturesBitShift(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "No flags set",
		},
		{
			name: "First flag set",
		},
		{
			name: "Second flag set",
		},
		{
			name: "First Two flags set",
		},
		{
			name: "Third flag set",
		},
		{
			name: "First and Third flag set",
		},
		{
			name: "Second and Third flag set",
		},
		{
			name: "First Three flags set",
		},
		{
			name: "Fourth flags set",
		},
		{
			name: "First and Fourth flag set",
		},
		{
			name: "First, Second and Fourth flag set",
		},
		{
			name: "First, Third and Fourth flag set",
		},
		{
			name: "Second, Third and Fourth flag set",
		},
		{
			name: "All Flags set",
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &CreItem{
				Flags: uint32(i),
			}
			out, err := item.MarshalJSON()
			if err != nil {
				t.Fatalf("failed to marshall")
			}
			expected := &CreItem{}
			if err := json.Unmarshal(out, expected); err != nil {
				t.Fatalf("failed to unmarshal")
			}
			if item.Flags != expected.Flags {
				t.Fatalf("Result:\n%d\n Does not match Expected:\n%d\n", item.Flags, expected.Flags)
			}
		})
	}
}
