package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct {
	game tortuga.Console
}

var (
	game  = tortuga.New()
	rect  = geom.MakeRect(float64(game.GetScreenWidth()/2-10), float64(game.GetScreenHeight()/2-10), 20, 20)
	rect1 = geom.MakeRect(float64(game.GetScreenWidth()/2-10), float64(game.GetScreenHeight()/2-10), 20, 20)
)

const (
	speed = 4
)

func (c cart) Update() {
	if c.game.IsDownPressed() {
		rect1[1] += speed
	}
	if c.game.IsUpPressed() {
		rect1[1] -= speed
	}
	if c.game.IsLeftPressed() {
		rect1[0] -= speed
	}
	if c.game.IsRightPressed() {
		rect1[0] += speed
	}
}

func (c cart) Render() {
	d := game.GetDisplay()
	game.Clear()
	rect.Draw(d, game.Color(3))

	geom.MakePoint(rect1.GetCenter()).Draw(d, game.RGBA(uint8(12)))
	if rect.IsAxisAlignedCollision(rect1) {
		rect1.Draw(d, game.Color(4))
		return
	}
	rect1.Draw(d, game.Color(2))
}

func main() {
	game.Run(cart{
		game: game,
	})
}
