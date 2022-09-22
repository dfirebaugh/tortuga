package geom

import (
	"image/color"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinydraw"
)

type Circle struct {
	X float64
	Y float64
	R float64
}

func MakeCircle(x, y, r float64) Circle {
	return Circle{
		X: x,
		Y: y,
		R: r,
	}
}

func (c Circle) Draw(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.Circle(d, int16(c.X), int16(c.Y), int16(c.R*2), color)
}

func (c Circle) Filled(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledCircle(d, int16(c.X), int16(c.Y), int16(c.R*2), color)
}
