package main

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Orb struct {
	PhysicsObject
	Color color.Color
}

func (o *Orb) Update() {
	o.PhysicsObject.Update()
}

func (o *Orb) Render() {
	c := geom.MakeCircle(o.X*pixelPerMeter, o.Y*pixelPerMeter, o.Height)
	geom.Fill(c, game.GetDisplay(), o.Color)
}
