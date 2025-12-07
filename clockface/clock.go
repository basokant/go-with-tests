package clockface

import (
	"math"
	"time"
)

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time 't'
// represented as a Point.
func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

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
