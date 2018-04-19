package srtm

import (
	"io"
	"strings"

	"github.com/pkg/errors"
)

var InvalidHGTFileName = errors.New("invalid HGT file name")

// TODO BAD! USE READER INTERFACE!? why tho

// ReadFile reads a SRTM data file (hgt)
// SRTM data must come from a file because
// file naming is part of the data format
func ReadFile(file string) (io.Reader, error) {
	corner, err := GetFileCorner(file)
	if err != nil {
		return nil, errors.Wrap(err, "could not get SRTM data file corner coords")
	}

	if strings.HasSuffix(file, ".gz") {

	}
	return nil, nil
}

func Read(r io.Reader) {

}

// GetFileCorner returns the southwest point contained in a HGT file.
// Coordinates in the file are relative to this point
func GetFileCorner(file string) (p Point, err error) {
	fnameParts := srtmParseName.FindStringSubmatch(file)
	if fnameParts != nil {
		return nil, InvalidHGTFileName
	}
	return p, err
}
