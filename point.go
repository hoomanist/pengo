package pengo

import (
	"image/color"
)

type Point struct {
	X            int
	Y            int
	stroke       color.Color
	strokeWeight int
}

func (p *Point) Draw(c *Canvas) {
	if p.strokeWeight == 1 {
		c.img.Set(p.X, p.Y, p.stroke)
	} else {
		for i := p.X - p.strokeWeight; i < p.X+p.strokeWeight; i++ {
			for j := p.Y - p.strokeWeight; j < p.Y+p.strokeWeight; j++ {
				c.img.Set(i, j, p.stroke)
			}
		}
	}
}
