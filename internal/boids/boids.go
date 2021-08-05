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
	Body                 *imdraw.IMDraw
	Width                float64
	Height               float64
	Origin               *tools.Coordinates
	Vertex               *tools.Coordinates
	Speed                uint16
	AngleSpeed           float64 // Angle speed is saved in radians
	Angle                float64 // Angle is saved in radians
	TravelAngle          float64
	TravelAngleDirection bool // true is left, false is right

	MaxTravelAngleValue uint16
	MaxTravelAngle      float64
}

func (b *Boid) construct(originX float64, originY float64, angle float64) {
	b.Width = 16
	b.Height = 20
	b.Speed = 100 // Defautl 180
	b.AngleSpeed = 20 * math.Pi / 180
	b.MaxTravelAngleValue = 13
	b.Origin = &tools.Coordinates{}
	b.Origin.X = originX
	b.Origin.Y = originY
	b.Angle = angle

	b.Body = imdraw.New(nil)
	b.Body.Color = pixel.RGB(0, 0, 0)

	//Find vertex
	b.Vertex = trig.FindPointFromPoint(b.Origin, b.TravelAngle, b.Height)
}

func New(x float64, y float64, angle float64, travelAngleDirection bool) *Boid {
	b := Boid{}

	// Construct
	b.construct(x, y, angle)

	if travelAngleDirection == true {
		b.TravelAngleDirection = travelAngleDirection
		b.MaxTravelAngle = b.Angle + float64(b.MaxTravelAngleValue)*math.Pi/180
	} else if travelAngleDirection == false {
		b.TravelAngleDirection = travelAngleDirection
		b.MaxTravelAngle = b.Angle - float64(b.MaxTravelAngleValue)*math.Pi/180
	}
	b.TravelAngle = b.Angle

	//Position
	b.drawBoidBody()
	return &b
}

func (b *Boid) GetSidePoints() (leftPoint, rightPoint *tools.Coordinates) {
	angleLeftSidePoint := (math.Pi / 2) + b.TravelAngle
	angleRightSidePoint := (math.Pi / 2) + math.Pi + b.TravelAngle
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

func Generator(boidsNumber int16) []*Boid {
	boids := make([]*Boid, boidsNumber)
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for i := 0; i < int(boidsNumber); i++ {
		x := rand.Intn(borders.WindowWidth-borders.BordersWidth*2) + borders.BordersWidth*2
		y := rand.Intn(borders.WindowHeight-borders.BordersWidth*2) + borders.BordersWidth*2

		angle := float64(rand.Intn(360))
		angleRad := angle * (math.Pi / 180)

		var travelAngleDirection bool
		if int(angle)%2 == 1 {
			travelAngleDirection = true
		} else {
			travelAngleDirection = false
		}
		boids[i] = New(float64(x), float64(y), angleRad, travelAngleDirection)
	}
	return boids
}

func (b *Boid) Live() {
	b.MoveForward()
}
