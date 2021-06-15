package trig

import (
	"math"

	"github.com/MydroX/goids/tools"
)

func FindPointFromPoint(point tools.Coordinates, angle float64, dist float64) *tools.Coordinates {
	newPoint := &tools.Coordinates{}

	if angle > math.Pi/2 && angle < (math.Pi/2)+math.Pi {
		newPoint.X = point.X - (math.Cos(angle) * dist)
	} else if angle == math.Pi/2 || angle == (math.Pi/2)+math.Pi {
		newPoint.X = point.X
	} else {
		newPoint.X = point.X + (math.Cos(angle) * dist)
	}

	if angle > math.Pi && angle < math.Pi*2 {
		newPoint.Y = point.Y - (math.Sin(angle) * dist)
	} else if angle == math.Pi || angle == math.Pi*2 || angle == 0 {
		newPoint.Y = point.Y
	} else {
		newPoint.Y = point.Y + (math.Sin(angle) * dist)
	}

	return newPoint
}
