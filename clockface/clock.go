package clockface

import (
	"math"
	"time"
)

type Point = struct {
	X float64
	Y float64
}

func secondsInRadians(tm time.Time) float64 {
	return (math.Pi / (30 / (float64(tm.Second()))))
}

func secondHandPoint(tm time.Time) Point {
	angle := secondsInRadians(tm)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
