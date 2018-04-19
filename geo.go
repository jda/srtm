package srtm

import (
	"log"
	"regexp"
	"strconv"
)

// should this be option to support non-30m data?
const squareSize = 3601 // 30m SRTM data is 3601 squares tall/wide

var srtmParseName = regexp.MustCompile(`(N|S)(\d\d)(E|W)(\d\d\d)\.hgt(\.gz)?`)
var negativeDD = regexp.MustCompile(`S|W`)

type Point struct {
	Latitude  float64
	Longitude float64
	Elevation float64
}

func dToDecimal(d string) (dd float64) {
	i, err := strconv.Atoi(d[1:])
	if err != nil {
		log.Fatalf("error converting Degree [%s] to Decimal: %s", d, err)
	}

	dd = float64(i)

	if negativeDD.MatchString(d) {
		dd = dd * -1
	}

	return dd
}
