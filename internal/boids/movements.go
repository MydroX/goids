package boids

import (
	"math"

	"github.com/MydroX/goids/internal/borders"
	"github.com/MydroX/goids/tools"
	"github.com/MydroX/goids/tools/trig"
)

func (b *Boid) MoveForward() {
	b.TravelRotation()

	isColliding, collidingPositon := b.IsCollidingBorder()
	if isColliding {
		b.Rotate(collidingPositon)
	}

	newOriginCoordinates := trig.FindPointFromPoint(b.Origin, b.TravelAngle, float64(b.Speed)*tools.DeltaTime)
	b.Origin = newOriginCoordinates

	newVertexCoordinates := trig.FindPointFromPoint(b.Origin, b.TravelAngle, b.Height)
	b.Vertex = newVertexCoordinates

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

	b.TravelAngle = b.Angle

	// MAY BE REFACTOR THIS PART FOR SMOOTHER ROTATION
	b.Origin = b.Vertex
	b.Vertex = trig.FindPointFromPoint(b.Origin, b.Angle, b.Height)
}

func (b *Boid) TravelRotation() {
	if b.TravelAngleDirection == true && b.TravelAngle >= b.MaxTravelAngle {
		b.TravelAngleDirection = false
		b.MaxTravelAngle = b.Angle - float64(b.MaxTravelAngleValue)*math.Pi/180
	} else if b.TravelAngleDirection == false && b.TravelAngle <= b.MaxTravelAngle {
		b.TravelAngleDirection = true
		b.MaxTravelAngle = b.Angle + float64(b.MaxTravelAngleValue)*math.Pi/180
	}

	if b.MaxTravelAngle > math.Pi*2 {
		b.TravelAngle = b.MaxTravelAngle - math.Pi*2
	}
	if b.MaxTravelAngle < 0 {
		b.Angle = b.MaxTravelAngle + math.Pi*2
	}

	if b.TravelAngleDirection == true {
		b.TravelAngle = b.TravelAngle + tools.DeltaTime*b.AngleSpeed
	} else if b.TravelAngleDirection == false {
		b.TravelAngle = b.TravelAngle - tools.DeltaTime*b.AngleSpeed
	}
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
