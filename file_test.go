package srtm

import (
	"os"
	"path"
	"testing"
)

var testFiles = []string{
	"S46W066.hgt.gz",
	"S46W067.hgt",
}

func TestReadFile(t *testing.T) {
	wd, _ := os.Getwd()
	for _, tf := range testFiles {
		t.Run(tf, func(t *testing.T) {
			tFileName := path.Join(wd, "testdata", tf)

			_, err := ReadFile(tFileName)
			if err != nil {
				t.Errorf("failed to process %s: %s", tFileName, err)
			}
		})
	}

}
