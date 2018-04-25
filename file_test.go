package srtm

import (
	"os"
	"path"
	"testing"
)

func TestReadFile(t *testing.T) {
	wd, _ := os.Getwd()
	tFileName := path.Join(wd, "testdata", "S49W066.hgt.gz")

	_, err := ReadFile(tFileName)
	if err != nil {
		t.Errorf("failed to process %s: %s", tFileName, err)
	}

}
