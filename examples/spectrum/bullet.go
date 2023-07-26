package main

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Bullet struct {
	PhysicsObject
	Thrust      float64
	Color       color.Color
	LifeTime    uint8
	MaxLifeTime uint8
}

func (b *Bullet) Update() {
	b.PhysicsObject.Update()
	b.LifeTime++
}

func (b *Bullet) Render() {
	c := geom.MakeCircle(b.X*pixelPerMeter, b.Y*pixelPerMeter, b.Height)
	geom.Fill(c, game.GetDisplay(), b.Color)
}
