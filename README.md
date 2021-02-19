# pengo
Go library for visual arts

## installation

``` go get -u  github.com/hoomanist/pengo ```

note: it is not that useful yet

### Getting Started

First you should create a new canvas

``` 
canvas := pengo.NewCanvas(400, 600, color.RGBA{0x00, 0xff, 0x00, 0xff})

```

this will create a 400x600 canvas with color green.

you can add a point on canvas

``` 
some_color := color.RGBA{0xff, 0x00, 0x00, 0xff}
point := pengo.Point{
    100, //x position
    200, //y position
    some_color, //color
    2, // thickness
}
point.Draw(&canvas)

```
 also a line:

 ``` 
 line := pengo.Line{
    0 , // start position
    600, // end position
    200, // line position
    "Vertical", // direction
    some_color, // color
 }
 line.Draw(&canvas)
 ```