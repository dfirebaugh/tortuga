package geom

import (
	"image/color"
	"math"

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

func (c Circle) Diameter() float64 {
	return 2 * c.R
}

func (c Circle) HasOverlap(other Circle) bool {
	return math.Pow((c.X-other.X), 2)+math.Pow((c.Y-other.Y), 2) <=
		math.Pow(2*(c.R+other.R), 2)
}

func (c Circle) ContainsPoint(p Point) bool {
	return math.Abs((c.X-p.X)*(c.X-p.X)+(c.Y-p.Y)*(c.Y-p.Y)) < ((c.Diameter()) * (c.Diameter()))
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
