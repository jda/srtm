# srtm
Go library for reading [Shuttle Radar Topography Mission](https://en.wikipedia.org/wiki/Shuttle_Radar_Topography_Mission) (SRTM) HGT files

[![GoDoc](https://godoc.org/github.com/jda/srtm?status.svg)](https://godoc.org/github.com/jda/srtm)
[![Go Report Card](https://goreportcard.com/badge/github.com/jda/srtm)](https://goreportcard.com/report/github.com/jda/srtm)
[![Build Status](https://travis-ci.org/jda/srtm.png)](https://travis-ci.org/jda/srtm)

```go
package main

import (
	"github.com/jda/srtm"
	"log"
)

func main() {
	geo, err := srtm.ReadFile("srtm/testdata/S46W067.hgt")
	if err != nil {
		log.Fatal(err)
	}
	p := geo[30]
	log.Printf("Lat: %.4f, Lng: %.4f, Elevation: %d", p.Latitude, p.Longitude, p.Elevation)
}
```

## Limitations
1. Until [issue #3](https://github.com/jda/srtm/issues/3) is resolved, only 1-arcsecond tiles are supported.
