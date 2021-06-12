package borders

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	Width  = 1200
	Height = 768
)

func Borders() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(240, 240, 240)
	// imd.Color = pixel.Alpha(0)
	imd.Push(pixel.V(15, 15))
	imd.Push(pixel.V(Width-15, Height-15))
	imd.Rectangle(0)
	return imd
}
