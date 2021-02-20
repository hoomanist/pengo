package pengo

import (
	"image/color"
)

type Rect struct {
	min   Point
	max   Point
	color color.Color
}

func (r *Rect) Draw(c *Canvas) {
	line0 := Line{
		r.min.X,
		r.max.X,
		r.min.Y,
		"Horizontal",
		r.color,
	}
	line0.Draw(c)
	line1 := Line{
		r.min.X,
		r.max.X,
		r.max.Y,
		"Horizontal",
		r.color,
	}
	line1.Draw(c)
	line2 := Line{
		r.min.Y,
		r.max.Y,
		r.min.X,
		"Horizontal",
		r.color,
	}
	line2.Draw(c)
	line3 := Line{
		r.min.Y,
		r.max.Y,
		r.max.Y,
		"Horizontal",
		r.color,
	}
	line3.Draw(c)
}
