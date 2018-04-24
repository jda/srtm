package srtm

import (
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// ErrInvalidHGTFileName is returned when a HGT file name does not match
// pattern of a valid file which indicates lat/lon
var ErrInvalidHGTFileName = errors.New("invalid HGT file name")

// ReadFile is a helper func around Read that reads a SRTM file, decompressing
// if necessary, and returns  SRTM elevation data
func ReadFile(file string) (points []Point, err error) {
	f, err := os.Open(file)
	if err != nil {
		return points, err
	}
	defer f.Close()

	if strings.HasSuffix(file, ".gz") {
		rdr, err := gzip.NewReader(f)
		if err != nil {
			return points, err
		}

		defer rdr.Close()
		return Read(file, rdr)
	}

	return Read(file, f)
}

// Read reads elevation for points from a SRTM file
func Read(fname string, r io.Reader) (points []Point, err error) {
	swCorner, err := GetFileCorner(fname)
	if err != nil {
		return points, errors.Wrap(err, "could not get corner coordinates from file name")
	}

	points = make([]Point, squareSize*squareSize)
	pIdx := 0

	// Latitude
	for row := 0; row < squareSize; row++ {
		lat := swCorner.Latitude + float64(row)/float64(squareSize)

		// Longitude
		for col := 0; col < squareSize; col++ {
			lon := swCorner.Longitude + float64(col)/float64(squareSize)

			var elev int16
			readErr := binary.Read(r, binary.BigEndian, &elev)
			if readErr != nil {
				return points, errors.Wrapf(err, "EOF before %d?", squareSize)
			}

			points[pIdx] = Point{
				Latitude:  lat,
				Longitude: lon,
				Elevation: elev,
			}
			pIdx++
		}
	}

	return points, nil
}

// GetFileCorner returns the southwest point contained in a HGT file.
// Coordinates in the file are relative to this point
func GetFileCorner(file string) (p Point, err error) {
	fnameParts := srtmParseName.FindStringSubmatch(file)
	if fnameParts == nil {
		return p, ErrInvalidHGTFileName
	}

	swLatitude, err := dToDecimal(fnameParts[1] + fnameParts[2])
	if err != nil {
		return p, errors.Wrap(err, "could not get Latitude from file name")
	}
	swLongitude, err := dToDecimal(fnameParts[3] + fnameParts[4])
	if err != nil {
		return p, errors.Wrap(err, "could not get Longitude from file name")
	}

	p = Point{
		Latitude:  swLatitude,
		Longitude: swLongitude,
	}

	return p, err
}
