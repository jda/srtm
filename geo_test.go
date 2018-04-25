package srtm

import (
	"testing"

	"github.com/pkg/errors"
)

func TestParseGeo(t *testing.T) {
	v := "N33"
	_, err := dToDecimal(v)
	if err != nil {
		t.Errorf("could not parse %s: %s", v, err)
	}
}

func TestParseInvalidCoordsize(t *testing.T) {
	v := "W181"
	res, err := dToDecimal(v)
	if errors.Cause(err) != ErrInvalidCoordDegrees {
		t.Errorf("parsing %s should throw error, instead got %s and value %f", v, err, res)
	}
}

func TestParseInvalidDirVals(t *testing.T) {
	v := "X10"
	res, err := dToDecimal(v)
	if errors.Cause(err) != ErrInvalidCoordDegrees {
		t.Errorf("parsing %s should throw error, instead got %s and value %f", v, err, res)
	}
}
