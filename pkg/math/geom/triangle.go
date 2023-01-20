package geom

import (
	"image/color"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinydraw"
)

type Triangle [3]Vector

func MakeTriangle(vecs [3]Vector) Triangle {
	return Triangle(vecs)
}

func (t Triangle) Draw(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.Triangle(d, int16(t[0][0]), int16(t[0][1]), int16(t[1][0]), int16(t[1][1]), int16(t[2][0]), int16(t[2][1]), color)
}

func (t Triangle) Filled(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledTriangle(d, int16(t[0][0]), int16(t[0][1]), int16(t[1][0]), int16(t[1][1]), int16(t[2][0]), int16(t[2][1]), color)
}
