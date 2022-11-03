package main

import (
	"flag"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/wad"
)

var (
	wadFile string
	tileKey string
)

// displays assets from a wad file for demo purposes
// e.g.
//
//	go run ./cmd/mapviewer/ -wad ./cmd/mapviewer/examples/game.wad -key tiles
func main() {
	flag.StringVar(&wadFile, "wad", "", "pass in a filepath to preview a map")
	flag.StringVar(&tileKey, "key", "", "key of tiles to render")
	flag.Parse()

	game := tortuga.New()
	if wadFile != "" {
		w := wad.New(wadFile)
		game.SetPalette(w.GetPalettes().Get("main"))
		game.SetTiles(w.GetTileDefinitions(), w.GetBackgrounds().Get(tileKey))
	}
	game.SetFPSEnabled(true)
	game.SetScaleFactor(3)
	game.Run(cart{})
}
