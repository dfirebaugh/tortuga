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
	game  tortuga.Console
	rect  geom.Rect
	speed = 4.0
)

func (c cart) Update() {
	if c.input.IsDownPressed() {
		rect[1] += speed
	}
	if c.input.IsUpPressed() {
		rect[1] -= speed
	}
	if c.input.IsLeftPressed() {
		rect[0] -= speed
	}
	if c.input.IsRightPressed() {
		rect[0] += speed
	}
}

func (c cart) Render() {
	game.Clear()

	// render a rectangle on the given display as a certain color
	// draw calls need to happen in the render loop
	rect.Draw(game.GetDisplay(), game.Color(2))
}

func main() {
	// create a rectangle when the app starts (so we don't create on every render loop)
	rect = geom.MakeRect(20, 20, 20, 20)

	// instantiate the game console
	game = tortuga.New()

	// run the cart (note: this is a blocking operation)
	game.Run(cart{
		input: input.Keyboard{},
	})
}
