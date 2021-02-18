package pengo

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// the canvas itseld
type Canvas struct {
	width   int
	height  int
	quality int
	img     *image.RGBA
}

// create a new canvas with 80% quality
func NewCanvas(width, height int, clr color.Color) Canvas {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, clr)
		}
	}
	canvas := Canvas{
		width:   width,
		height:  height,
		quality: 80,
		img:     img,
	}
	return canvas
}

// change default quality from 1 to 100
func (c *Canvas) SetQuality(quality int) {
	c.quality = quality
}

// save canvas into a file
func (c *Canvas) Save(filename string, filetype string) error {
	f, err := os.Create("filename")
	if err != nil {
		return err
	}
	if filetype == "PNG" {
		err = png.Encode(f, c.img)
	} else if filetype == "JPG" {
		err = jpeg.Encode(f, c.img, &jpeg.Options{
			Quality: c.quality,
		})
	}
	return err
}

type Shape interface {
	Draw(c *Canvas)
}
