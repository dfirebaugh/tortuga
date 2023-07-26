package geom

import (
	"image/color"
	"math"

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

func (t *Triangle) Centroid() Point {
	return Point{
		X: (t[0][0] + t[1][0] + t[2][0]) / 3,
		Y: (t[0][1] + t[1][1] + t[2][1]) / 3,
	}
}

func (t *Triangle) Rotate(angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	centroid := t.Centroid()

	// Translate points to origin
	t[0][0] -= centroid.ToVector()[0]
	t[0][1] -= centroid.ToVector()[1]
	t[1][0] -= centroid.ToVector()[0]
	t[1][1] -= centroid.ToVector()[1]
	t[2][0] -= centroid.ToVector()[0]
	t[2][1] -= centroid.ToVector()[1]

	// Rotate points around the origin
	t[0][0], t[0][1] = cos*t[0][0]-sin*t[0][1], sin*t[0][0]+cos*t[0][1]
	t[1][0], t[1][1] = cos*t[1][0]-sin*t[1][1], sin*t[1][0]+cos*t[1][1]
	t[2][0], t[2][1] = cos*t[2][0]-sin*t[2][1], sin*t[2][0]+cos*t[2][1]

	// Translate points back to the original position
	t[0][0] += centroid.X
	t[0][1] += centroid.Y
	t[1][0] += centroid.X
	t[1][1] += centroid.Y
	t[2][0] += centroid.X
	t[2][1] += centroid.Y
}
