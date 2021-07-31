package borders

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	WindowWidth  = 1200
	WindowHeight = 768
	BordersWidth = 15
)

type Borders struct {
	Width float64
}

func New() Borders {
	return Borders{
		Width: BordersWidth,
	}
}

func (b *Borders) Draw() *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(240, 240, 240)
	imd.Push(pixel.V(b.Width, b.Width))
	imd.Push(pixel.V(WindowWidth-b.Width, WindowHeight-b.Width))
	imd.Rectangle(0)
	return imd
}
