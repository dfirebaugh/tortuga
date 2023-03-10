package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct{}

func (c cart) Update() {
}

func (c cart) Render() {
	geom.MakeRect(0, 0, float64(game.GetScreenWidth()), float64(game.GetScreenHeight())).Filled(game.GetDisplay(), game.Color(2))
}

var game tortuga.Console

func blockFactory(v uint8) []uint8 {
	var b []uint8
	for i := 0; i < 8*8; i++ {
		b = append(b, v)
	}
	return b
}

func main() {
	game = tortuga.New()
	for i := 0; i < game.GetScreenWidth()/game.GetTileSize(); i++ {
		game.SetTile(i, 0, blockFactory(uint8(i+1)))
		game.SetTile(0, i, blockFactory(uint8(i+1)))
		game.SetTile(i, i, blockFactory(uint8(i+1)))
	}
	game.SetTile(2, 5, blockFactory(7))

	game.SetFPSEnabled(true)
	game.SetScaleFactor(5)
	game.Run(cart{})
}
