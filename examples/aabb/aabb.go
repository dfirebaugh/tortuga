package main

import (
	"tortuga/pkg/input"
	"tortuga/pkg/math/geom"
	"tortuga/pkg/tortuga"
)

type cart struct{}

var (
	game  = tortuga.New()
	i     = input.Keyboard{}
	rect  = geom.MakeRect(float64(game.GetScreenWidth()/2-10), float64(game.GetScreenHeight()/2-10), 20, 20)
	rect1 = geom.MakeRect(float64(game.GetScreenWidth()/2-10), float64(game.GetScreenHeight()/2-10), 20, 20)
)

const (
	speed = 4
)

func (c cart) Update() {
	if i.IsDownPressed() {
		rect1[1] += speed
	}
	if i.IsUpPressed() {
		rect1[1] -= speed
	}
	if i.IsLeftPressed() {
		rect1[0] -= speed
	}
	if i.IsRightPressed() {
		rect1[0] += speed
	}
}

func (c cart) Render() {
	game.Clear()
	rect.Render(game.GetDisplay(), game.Color(3))

	if rect.IsAxisAlignedCollision(rect1) {
		rect1.Render(game.GetDisplay(), game.Color(4))
		return
	}
	rect1.Render(game.GetDisplay(), game.Color(2))

	geom.MakePoint(rect1.GetCenter()).Render(game.GetDisplay(), uint8(12))
}

func main() {
	game.Run(cart{})
}
