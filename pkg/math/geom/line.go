package geom

import (
	"errors"
	"image/color"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinydraw"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Draw(d displayer, clr uint8) {
	d.Put(int16(p.X), int16(p.Y), clr)
}

type Segment struct {
	v0 Vector
	v1 Vector
}

type Line struct {
	slope   float64
	yint    float64
	segment Segment
}

func MakePoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func MakeLine(a, b Point) Line {
	slope := (b.Y - a.Y) / (b.X - a.X)
	yint := a.Y - slope*a.X
	return Line{
		slope: slope,
		yint:  yint,
		segment: Segment{
			v0: MakeVector(a.X, a.Y),
			v1: MakeVector(b.X, b.Y),
		},
	}
}

func (l Line) EvalX(x float64) float64 {
	return l.slope*x + l.yint
}

func (l Line) IsParrallel(l1, l2 Line) bool {
	return l1.slope == l2.slope
}

func (l Line) Intersection(l2 Line) (Point, error) {
	if l.slope == l2.slope {
		return Point{}, errors.New("the lines do not intersect")
	}
	x := (l2.yint - l.yint) / (l.slope - l2.slope)
	y := l.EvalX(x)

	return Point{x, y}, nil
}

func (l Line) Draw(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.Line(d, int16(l.segment.v0[0]), int16(l.segment.v0[1]), int16(l.segment.v1[0]), int16(l.segment.v1[1]), color)
}
