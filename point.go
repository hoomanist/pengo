package pengo

import (
	"image/color"
)

type Point struct {
	X      int
	Y      int
	stroke color.Color
}

func (p *Point) Draw(c *Canvas) {
	c.img.Set(p.X, p.Y, p.stroke)
}
