package trig

import (
	"math"

	"github.com/MydroX/goids/tools"
)

func FindPointFromPoint(originPoint *tools.Coordinates, angle float64, dist float64) *tools.Coordinates {
	newPoint := &tools.Coordinates{}

	sin, cos := math.Sincos(angle)

	newPoint.X = (cos * dist) + originPoint.X
	newPoint.Y = (sin * dist) + originPoint.Y

	return newPoint
}
