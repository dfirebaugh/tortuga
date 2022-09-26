package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/input"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct {
	input input.PlayerInput
}

var (
	game = tortuga.New()
	rect = geom.MakeRect(float64(game.GetScreenWidth()/2-20), float64(game.GetScreenHeight()/2-20), 40, 40)
	ray  = geom.Ray{
		Origin:    geom.MakeVector(60, 60),
		Direction: geom.MakeVector(input.CursorPositionFloat()),
	}
	collision = &geom.Collision{}
)

func (c cart) Render() {
	d := game.GetDisplay()
	game.Clear()
	renderRay()

	if rect.HasRayIntersection(ray, collision) {
		rect.Filled(d, game.Color(4))

		// draw the collision point
		geom.MakeCircle(collision.Point[0], collision.Point[1], 2).Filled(d, game.Color(3))

		// draw the collision normal
		normal := collision.Point.Add(collision.Normal.Multiply(geom.MakeVector(10, 10))).ToPoint()
		geom.MakeLine(
			collision.Point.ToPoint(),
			normal).Draw(d, game.Color(3))

		return
	}

	rect.Draw(d, game.Color(3))
}

func (c cart) Update() {
	ray.Direction = geom.MakeVector(input.CursorPositionFloat())
	if c.input.IsDownPressed() {
		rect[1] += 5
	}
	if c.input.IsUpPressed() {
		rect[1] -= 5
	}
	if c.input.IsLeftPressed() {
		rect[0] -= 5
	}
	if c.input.IsRightPressed() {
		rect[0] += 5
	}
}

func renderDirection() {
	d := game.GetDisplay()
	geom.MakeCircle(rect[0], rect[1], 1).Filled(d, game.Color(1))
}

func renderRay() {
	d := game.GetDisplay()
	ray.Draw(d, game.Color(5))

	renderDirection()
}

func main() {
	game.Run(cart{
		input: input.Keyboard{},
	})
}
