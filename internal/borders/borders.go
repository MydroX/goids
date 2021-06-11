package borders

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func Borders() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(240, 240, 240)
	// imd.Color = pixel.Alpha(0)
	imd.Push(pixel.V(12, 12))
	imd.Push(pixel.V(1012, 754))
	imd.Rectangle(0)
	return imd
}
