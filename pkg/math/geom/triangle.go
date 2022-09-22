package geom

import (
	"image/color"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinydraw"
)

type triangle [3]Vector

func MakeTriangle(vecs [3]Vector) triangle {
	return triangle(vecs)
}

func (t triangle) Draw(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.Triangle(d, int16(t[0][0]), int16(t[0][1]), int16(t[1][0]), int16(t[1][1]), int16(t[2][0]), int16(t[2][1]), color)
}

func (t triangle) Filled(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledTriangle(d, int16(t[0][0]), int16(t[0][1]), int16(t[1][0]), int16(t[1][1]), int16(t[2][0]), int16(t[2][1]), color)
}
