package srtm

import (
	"errors"
	"regexp"
	"strconv"
)

// should this be option to support non-30m data?
const squareSize = 3601 // 30m SRTM data is 3601 squares tall/wide

var srtmParseName = regexp.MustCompile(`(N|S)(\d\d)(E|W)(\d\d\d)\.hgt(\.gz)?`)
var negativeDD = regexp.MustCompile(`S|W`)

var ErrInvalidCoordDegrees = errors.New("invalid lat/lon degrees")

type Point struct {
	Latitude  float64
	Longitude float64
	Elevation int16
}

func dToDecimal(d string) (dd float64, err error) {
	i, err := strconv.Atoi(d[1:])
	if err != nil {
		return dd, ErrInvalidCoordDegrees
	}

	dd = float64(i)

	if negativeDD.MatchString(d) {
		dd = dd * -1
	}

	return dd, nil
}
