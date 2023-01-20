package geom

import "image/color"

type Point struct {
	X float64
	Y float64
}

func MakePoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (p Point) ToVector() Vector {
	return MakeVector(p.X, p.Y)
}

func (p Point) Draw(d displayer, c color.RGBA) {
	d.SetPixel(int16(p.X), int16(p.Y), c)
}
