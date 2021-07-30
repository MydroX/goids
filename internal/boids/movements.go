package boids

import (
	"github.com/MydroX/goids/tools"
	"github.com/MydroX/goids/tools/trig"
)

func (b *Boid) MoveForward() {
	// TODO: Add border collision detection

	newVertexCoordinates := trig.FindPointFromPoint(b.Vertex, b.Angle, float64(b.Speed)*tools.DeltaTime)
	newOriginCoordinates := trig.FindPointFromPoint(b.Origin, b.Angle, float64(b.Speed)*tools.DeltaTime)

	b.Vertex = newVertexCoordinates
	b.Origin = newOriginCoordinates

	b.drawBoidBody()
}
