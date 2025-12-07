package clockface

import (
	"math"
	"time"
)

type Point = struct {
	X float64
	Y float64
}

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock

	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock

	hoursInHalfClock = 6
	hoursInClock     = 2 * hoursInHalfClock
)

func secondsInRadians(tm time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(tm.Second()))))
}

func secondHandPoint(tm time.Time) Point {
	return angleToPoint(secondsInRadians(tm))
}

func minutesInRadians(tm time.Time) float64 {
	secondsRadians := secondsInRadians(tm) / secondsInClock
	return secondsRadians + (math.Pi / (minutesInHalfClock / float64(tm.Minute())))
}

func minuteHandPoint(tm time.Time) Point {
	return angleToPoint(minutesInRadians(tm))
}

func hoursInRadians(tm time.Time) float64 {
	minutesRadians := minutesInRadians(tm) / minutesInClock
	return minutesRadians + (math.Pi / (hoursInHalfClock / float64(tm.Hour()%hoursInClock)))
}

func hourHandPoint(tm time.Time) Point {
	return angleToPoint(hoursInRadians(tm))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
