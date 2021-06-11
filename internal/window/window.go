package window

import (
	"time"

	"github.com/MydroX/goids/internal/borders"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Goids",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fps := time.NewTicker(time.Second / 60)
	for !win.Closed() {
		win.Clear(colornames.Grey)
		borders.Borders().Draw(win)
		win.Update()
		<-fps.C
	}
}
