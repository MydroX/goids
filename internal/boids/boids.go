package boids

import (
	"math"
	"math/rand"
	"time"

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
	Origin          *tools.Coordinates
	Vertex          *tools.Coordinates
	Angle           float64 // Angle are saved in radians
	Speed           uint8
}

func (b *Boid) construct(originX float64, originY float64, angle float64) {
	b.Width = 16
	b.Height = 20
	b.Speed = 80

	b.Origin = &tools.Coordinates{}
	b.Origin.X = originX
	b.Origin.Y = originY
	b.Angle = angle

	b.Body = imdraw.New(nil)
	b.Body.Color = pixel.RGB(0, 0, 0)

	//Find vertex
	b.Vertex = trig.FindPointFromPoint(b.Origin, b.Angle, b.Height)
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
	b.drawBoidBody()
	return b
}

func (b *Boid) GetSidePoints() (leftPoint, rightPoint *tools.Coordinates) {
	angleLeftSidePoint := (math.Pi / 2) + b.Angle
	angleRightSidePoint := (math.Pi / 2) + math.Pi + b.Angle
	return trig.FindPointFromPoint(b.Origin, angleLeftSidePoint, b.Width/2), trig.FindPointFromPoint(b.Origin, angleRightSidePoint, b.Width/2)
}

func (b *Boid) drawBoidBody() {
	b.Body.Clear()

	leftSidePoint, rightSidePoint := b.GetSidePoints()

	b.Body.Push(pixel.V(b.Vertex.X, b.Vertex.Y), pixel.V(leftSidePoint.X, leftSidePoint.Y))
	b.Body.Push(pixel.V(b.Vertex.X, b.Vertex.Y), pixel.V(rightSidePoint.X, rightSidePoint.Y))
	b.Body.Push(pixel.V(leftSidePoint.X, leftSidePoint.Y), pixel.V(rightSidePoint.X, rightSidePoint.Y))
	b.Body.Polygon(0)
}

// TO REWORK
//
// func (b *Boid) IsCollidingBorder() bool {
// 	//Get triangle vertex
// 	vertex := tools.Coordinates{
// 		X: b.Origin.X + b.Width,
// 		Y: b.Origin.Y + (b.Height / 2),
// 	}

// 	if vertex.X > borders.WindowWidth-borders.BordersWidth || vertex.X < borders.BordersWidth {
// 		return true
// 	}
// 	return false
// }

func Generator(boidsNumber int16) []Boid {
	boids := make([]Boid, boidsNumber)
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for i := 0; i < int(boidsNumber); i++ {
		x := rand.Intn(borders.WindowWidth-borders.BordersWidth*2) + borders.BordersWidth*2
		y := rand.Intn(borders.WindowHeight-borders.BordersWidth*2) + borders.BordersWidth*2

		angle := float64(rand.Intn(360))
		angleRad := angle * (math.Pi / 180)

		boids[i] = New(float64(x), float64(y), angleRad)
	}
	return boids
}
