package pengo

import (
	"image/color"
)

type Line struct {
	X1, X2, Y int
	Direction string
	Color     color.Color
}

func (l *Line) Draw(c *Canvas) {
	if l.Direction == "Horizontal" {
		for ; l.X1 <= l.X2; l.X1++ {
			c.img.Set(l.X1, l.Y, l.Color)
		}
	} else if l.Direction == "Vertical" {
		for ; l.X1 <= l.X2; l.X1++ {
			c.img.Set(l.Y, l.X1, l.Color)
		}
	}

}
