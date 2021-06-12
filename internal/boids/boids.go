package boids

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Vector struct {
	X float64
	Y float64
}

type Boid struct {
	Matrix *pixel.Matrix
	Draw   *imdraw.IMDraw
	Width  float64
	Height float64
	X      float64
	Y      float64
}

func New(x float64, y float64) Boid {
	b := Boid{}

	// Construct
	b.construct(x, y)

	//Position
	b.drawBoid(b.X, b.Y)

	return b
}

func (b *Boid) construct(x float64, y float64) {
	b.Width = 18
	b.Height = 14
	b.X = x
	b.Y = y
	b.Draw = imdraw.New(nil)
	b.Matrix = &pixel.Matrix{}

	b.Draw.Color = pixel.RGB(0, 0, 0)
}

func (b *Boid) drawBoid(x float64, y float64) {
	b.Draw.Clear()
	b.Draw.Push(pixel.V(x, y))
	b.Draw.Push(pixel.V(x-b.Width, y+(-b.Height/2)))
	b.Draw.Push(pixel.V(x-b.Width, y+b.Height/2))
	b.Draw.Polygon(0)
}

func (b *Boid) Move(v Vector) {
	b.X = b.X + v.X
	b.Y = b.Y + v.Y
	b.drawBoid(b.X, b.Y)
}
