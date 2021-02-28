package pengo

import (
	"image/color"
)

type Point struct {
	X, Y, StrokeWeight int
	Stroke             color.Color
}

func (p *Point) Draw(c *Canvas) {
	if p.StrokeWeight == 1 {
		c.img.Set(p.X, p.Y, p.Stroke)
	} else {
		for i := p.X - p.StrokeWeight; i < p.X+p.StrokeWeight; i++ {
			for j := p.Y - p.StrokeWeight; j < p.Y+p.StrokeWeight; j++ {
				c.img.Set(i, j, p.Stroke)
			}
		}
	}
}
