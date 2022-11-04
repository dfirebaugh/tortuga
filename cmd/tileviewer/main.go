package main

import (
	"flag"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/tile"
)

type cart struct{}

func (c cart) Update() {}
func (c cart) Render() {}

var (
	tileString string
	tileWidth  int
)

// displays a tile
// e.g.
//
//	go run ./cmd/tileviewer/ -tile 0880088088888888888887888888878888888888088888800088880000088000
func main() {
	flag.StringVar(&tileString, "tile", "", "encoded tile string")
	flag.IntVar(&tileWidth, "width", 8, "encoded tile string")
	flag.Parse()

	game := tortuga.New()
	if tileString != "" {
		t := tile.Decode(tileString)
		t.PixelSize = float64(game.GetScreenHeight() / tileWidth)
		t.Draw(game.GetDisplay())
	}
	game.SetScaleFactor(3)
	game.Run(cart{})
}
