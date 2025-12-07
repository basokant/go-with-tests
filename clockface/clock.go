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
	return angleToPoint(secondsInRadians(tm))
}

func minutesInRadians(tm time.Time) float64 {
	secondsRadians := secondsInRadians(tm) / 60
	return secondsRadians + (math.Pi / (30 / float64(tm.Minute())))
}

func minuteHandPoint(tm time.Time) Point {
	return angleToPoint(minutesInRadians(tm))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
