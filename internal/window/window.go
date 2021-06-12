package window

import (
	"time"

	"github.com/MydroX/goids/internal/boids"
	"github.com/MydroX/goids/internal/borders"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Goids",
		Bounds: pixel.R(0, 0, borders.Width, borders.Height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	canvas := win.Canvas()

	boidsXCoordinates := []float64{100}
	boidsYCoordinates := []float64{100}

	boid := boids.New(boidsXCoordinates[0], boidsYCoordinates[0])

	lastTime := time.Now()
	fps := time.NewTicker(time.Second / 60)
	for !win.Closed() {
		//Delta time
		dt := time.Since(lastTime).Seconds()
		lastTime = time.Now()

		canvas.Clear(colornames.Grey)

		// Draw borders
		borders.Borders().Draw(canvas)

		// Draw boids
		boid.Draw.Draw(canvas)

		// Move 1st boid
		boid.Move(boids.Vector{X: 100 * dt, Y: 0})

		win.Update()
		<-fps.C
	}
}
