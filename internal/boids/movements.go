package boids

import (
	"math"

	"github.com/MydroX/goids/internal/borders"
	"github.com/MydroX/goids/tools"
	"github.com/MydroX/goids/tools/trig"
)

func (b *Boid) MoveForward() {
	isColliding, collidingPositon := b.IsCollidingBorder()
	if isColliding {
		b.Rotate(collidingPositon)
	}

	newVertexCoordinates := trig.FindPointFromPoint(b.Vertex, b.Angle, float64(b.Speed)*tools.DeltaTime)
	newOriginCoordinates := trig.FindPointFromPoint(b.Origin, b.Angle, float64(b.Speed)*tools.DeltaTime)

	b.Vertex = newVertexCoordinates
	b.Origin = newOriginCoordinates

	b.drawBoidBody()
}

func (b *Boid) Rotate(collidingPositon []int8) {
	if collidingPositon[0] == -1 {
		if b.Angle > math.Pi/2 && b.Angle < math.Pi {
			b.Angle = b.Angle - math.Pi/2
		} else if b.Angle > math.Pi && b.Angle < math.Pi+math.Pi/2 {
			b.Angle = b.Angle + math.Pi/2
		}
	} else if collidingPositon[0] == 1 {
		if b.Angle > math.Pi+math.Pi/2 && b.Angle < math.Pi*2 {
			b.Angle = b.Angle - math.Pi/2
		} else if b.Angle > 0 && b.Angle < math.Pi/2 {
			b.Angle = b.Angle + math.Pi/2
		}
	} else if collidingPositon[1] == -1 {
		if b.Angle > math.Pi && b.Angle < math.Pi+math.Pi/2 {
			b.Angle = b.Angle - math.Pi/2
		} else if b.Angle > math.Pi+math.Pi/2 && b.Angle < math.Pi*2 {
			b.Angle = b.Angle + math.Pi/2
		}
	} else if collidingPositon[1] == 1 {
		if b.Angle > 0 && b.Angle < math.Pi/2 {
			b.Angle = (b.Angle - math.Pi/2) + math.Pi*2
		} else if b.Angle > math.Pi/2 && b.Angle < math.Pi {
			b.Angle = b.Angle + math.Pi/2
		}
	}

	if b.Angle > math.Pi*2 {
		b.Angle = b.Angle - math.Pi*2
	}

	// MAY BE REFACTOR THIS PART FOR SMOOTHER ROTATION
	b.Origin = b.Vertex
	b.Vertex = trig.FindPointFromPoint(b.Origin, b.Angle, b.Height)
}

func (b *Boid) IsCollidingBorder() (bool, []int8) {
	if b.Vertex.X > borders.WindowWidth-borders.BordersWidth {
		return true, []int8{1, 0}
	} else if b.Vertex.X < borders.BordersWidth {
		return true, []int8{-1, 0}
	} else if b.Vertex.Y > borders.WindowHeight-borders.BordersWidth {
		return true, []int8{0, 1}
	} else if b.Vertex.Y < borders.BordersWidth {
		return true, []int8{0, -1}
	}
	return false, []int8{0, 0}
}
