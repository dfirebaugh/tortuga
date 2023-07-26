package main

import (
	"fmt"
	"image/color"

	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type ShipColor uint

const (
	red ShipColor = iota
	green
	blue
)

type Ship struct {
	PhysicsObject
	Color      color.Color
	lumenosity map[ShipColor]float32
}

const (
	rotationOffset = 1.6
)

func (s *Ship) Update() {
	s.PhysicsObject.Update()
}

func (s *Ship) Render() {
	x := s.X * pixelPerMeter
	y := s.Y * pixelPerMeter
	width := s.Width * pixelPerMeter
	height := s.Height * pixelPerMeter

	t := geom.MakeTriangle([3]geom.Vector{
		geom.MakeVector(x-width, y+height),
		geom.MakeVector(x+width, y+height),
		geom.MakeVector(x, y-height),
	})

	t.Rotate(float64(s.Rotation + rotationOffset))

	s.renderMeter()

	geom.Fill(t, game.GetDisplay(), s.Color)
	geom.Draw(t, game.GetDisplay(), game.Color(5))
}

func (s *Ship) renderMeter() {
	game.PrintAt(fmt.Sprintf("%d", int(s.lumenosity[red]*100)), 10, 20, 8)
	game.PrintAt(fmt.Sprintf("%d", int(s.lumenosity[green]*100)), 10, 30, 12)
	game.PrintAt(fmt.Sprintf("%d", int(s.lumenosity[blue]*100)), 10, 40, 11)
}
