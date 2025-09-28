package bg

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"testing"
)

const FixturesDirectory string = "../fixtures"

var (
	_, b, _, _  = runtime.Caller(0)
	basepath    = filepath.Dir(b)
	arefixtures = filepath.Join(basepath, FixturesDirectory, "are")
)

func TestReadWriteArea(t *testing.T) {
	err := filepath.WalkDir(arefixtures,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() || filepath.Ext(d.Name()) == ".json" {
				return nil
			}
			t.Run(d.Name(), func(t *testing.T) {
				if err != nil {
					t.Fatal(err)
				}
				file, err := os.Open(path)
				if err != nil {
					t.Fatal(err)
				}
				defer file.Close()
				area, err := OpenArea(file)
				if err != nil {
					t.Fatal(err)
				}
				if area == nil {
					t.Fatalf("Area is empty")
				}
				tempFile, err := os.CreateTemp("", "")
				if err != nil {
					t.Fatal(err)
				}
				defer tempFile.Close()
				if err := area.Write(tempFile); err != nil {
					t.Fatal(err)
				}
				tempFileContents := []byte{}
				if _, err := tempFile.Read(tempFileContents); err != nil {
					t.Fatal(err)
				}
				fixtureFileContents := []byte{}
				if _, err := file.Read(fixtureFileContents); err != nil {
					t.Fatal(err)
				}
				if !slices.Equal(fixtureFileContents, tempFileContents) {
					fmt.Printf("%v != %v\n", fixtureFileContents, tempFileContents)
					t.Fatalf("Binary contents of fixture do not match writen file")
				}
				fixture, err := os.ReadFile(path + ".json")
				if err != nil {
					t.Fatal(err)
				}
				expected := Area{}
				if err = json.Unmarshal(fixture, &expected); err != nil {
					t.Fatal(err)
				}
				if !area.Equal(&expected) {
					t.Fatalf("Result:\n%+v\n Does not match Expected:\n%+v\n", area, &expected)
				}
			})
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse area files, %+v", err)
	}
}
