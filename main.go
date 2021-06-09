package main

import (
	"github.com/MydroX/goids/internal/window"
	"github.com/faiface/pixel/pixelgl"
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawRectangle(0, 0, 0, 100)
	dc.Fill()

	pixelgl.Run(window.Run)
}
