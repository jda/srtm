package srtm

import "math"

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}

	return math.Floor(a + 0.5)
}
