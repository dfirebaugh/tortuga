package component

import (
	"fmt"

	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Coordinate struct {
	// x and y represents coordinates on the screen
	X float64
	Y float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d, %d", int(c.X), int(c.Y))
}

func (c *Coordinate) SetCoordinate(newCoord Coordinate) {
	c.X = newCoord.X
	c.Y = newCoord.Y
}

func (c Coordinate) GetDistance(other Coordinate) float64 {
	a := geom.MakeVector(c.X, c.Y)
	b := geom.MakeVector(other.X, other.Y)

	return a.GetDistance(b)
}

func (c Coordinate) GetDirection(other Coordinate) float64 {
	a := geom.MakeVector(c.X, c.Y)
	b := geom.MakeVector(other.X, other.Y)
	// point := a.Add(b).ToPoint()

	return a.GetDirection(b)
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

func (c Coordinate) Subtract(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X - other.X,
		Y: c.Y - other.Y,
	}
}

func (c Coordinate) TranslateXY(offset Coordinate, pixelSize float64) (float64, float64) {
	x := (c.X - offset.X) / pixelSize
	y := (c.Y - offset.Y) / pixelSize
	return x, y
}
