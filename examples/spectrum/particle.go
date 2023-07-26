package main

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Particle struct {
	PhysicsObject
	Thrust      float64
	Color       color.Color
	LifeTime    uint8
	MaxLifeTime uint8
}

func (p *Particle) Update() {
	p.PhysicsObject.Update()
	p.LifeTime++
}

func (p *Particle) Render() {
	c := geom.MakeCircle(p.X*pixelPerMeter, p.Y*pixelPerMeter, p.Height)
	geom.Fill(c, game.GetDisplay(), p.Color)
}
