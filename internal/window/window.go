package window

import (
	"time"

	"github.com/MydroX/goids/internal/boids"
	"github.com/MydroX/goids/internal/borders"
	"github.com/MydroX/goids/tools"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Goids",
		Bounds: pixel.R(0, 0, borders.WindowWidth, borders.WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	canvas := win.Canvas()

	borders := borders.New()

	boids := boids.Generator(1)
	boid := boids[0]

	lastTime := time.Now()
	fps := time.NewTicker(time.Second / 60)

	for !win.Closed() {
		//Delta time
		tools.DeltaTime = time.Since(lastTime).Seconds()
		lastTime = time.Now()

		canvas.Clear(colornames.Grey)

		// Draw borders
		borders.Draw().Draw(canvas)

		// Draw boids
		boid.Body.Draw(canvas)

		// Move 1st boid
		boid.MoveForward()

		win.Update()
		// win.SetClosed(true)
		<-fps.C
	}
}
