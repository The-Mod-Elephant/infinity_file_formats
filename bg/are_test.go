package bg

import (
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

func TestReadArea(t *testing.T) {
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
			defer file.Close()
			area, err := OpenArea(file)
			if err != nil {
				return err
			}
			if area == nil {
				return fmt.Errorf("Area nil failed")
			}
			tempFile, err := os.CreateTemp("", "")
			if err != nil {
				return err
			}
			defer tempFile.Close()
			if err := area.Write(tempFile); err != nil {
				return err
			}
			tempFileContents := []byte{}
			if _, err := tempFile.Read(tempFileContents); err != nil {
				return err
			}
			fixtureFileContents := []byte{}
			if _, err := file.Read(fixtureFileContents); err != nil {
				return err
			}
			if !slices.Equal(fixtureFileContents, tempFileContents) {
				fmt.Printf("%v != %v\n", fixtureFileContents, tempFileContents)
				return fmt.Errorf("Binary contents of fixture do not match writen file")
			}
			return nil
		})
	if err != nil {
		t.Fatalf("Failed to parse area files, %+v", err)
	}
}
