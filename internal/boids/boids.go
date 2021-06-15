package boids

import (
	"math/rand"

	"github.com/MydroX/goids/internal/borders"
	"github.com/MydroX/goids/tools"
	"github.com/MydroX/goids/tools/trig"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Boid struct {
	Body            *imdraw.IMDraw
	MovingDirection *tools.Vector
	Width           float64
	Height          float64
	X               float64
	Y               float64
	Vertex          *tools.Coordinates
	Angle           float64 // Angle are saved in degrees
}

func New(x float64, y float64, angle float64) Boid {
	b := Boid{}

	// Construct
	b.construct(x, y, angle)

	//Initial movement
	b.MovingDirection = &tools.Vector{
		X: 100,
		Y: 0,
	}

	//Position
	b.drawBoidBody(b.X, b.Y)
	return b
}

func (b *Boid) construct(originX float64, originY float64, angle float64) {
	b.Width = 18
	b.Height = 14

	b.X = originX
	b.Y = originY
	b.Angle = angle

	b.Angle = 0

	b.Body = imdraw.New(nil)
	b.Body.Color = pixel.RGB(0, 0, 0)

	//Find vertex
	boidOrigin := tools.Coordinates{X: b.X, Y: b.Y}
	b.Vertex = trig.FindPointFromPoint(boidOrigin, b.Angle, b.Height)
}

func (b *Boid) drawBoidBody(x float64, y float64) {
	b.Body.Clear()

	b.Body.Push(pixel.V(x, y))
	b.Body.Push(pixel.V(x, y+b.Height))
	b.Body.Push(pixel.V(x+b.Width, y+b.Height/2))
	b.Body.Polygon(0)
}

func (b *Boid) Move() {
	if b.IsCollidingBorder() {
		b.MovingDirection.X = -b.MovingDirection.X
	}

	correctedSpeedX := b.MovingDirection.X * tools.DeltaTime
	correctedSpeedY := b.MovingDirection.Y * tools.DeltaTime

	b.X = b.X + correctedSpeedX
	b.Y = b.Y + correctedSpeedY
	b.drawBoidBody(b.X, b.Y)
}

func (b *Boid) IsCollidingBorder() bool {
	//Get triangle vertex
	vertex := tools.Coordinates{
		X: b.X + b.Width,
		Y: b.Y + (b.Height / 2),
	}
	//
	if vertex.X > borders.WindowWidth-borders.BordersWidth || vertex.X < borders.BordersWidth {
		return true
	} else {
		return false
	}
}

func Generator(boidsNumber int16) []Boid {
	boids := make([]Boid, boidsNumber)

	for i := 0; i < int(boidsNumber); i++ {
		x := rand.Intn(borders.WindowWidth-borders.BordersWidth*2) + borders.BordersWidth*2
		y := rand.Intn(borders.WindowHeight-borders.BordersWidth*2) + borders.BordersWidth*2
		angle := float64(rand.Intn(360))

		boids[i] = New(float64(x), float64(y), angle)
	}
	return boids
}
