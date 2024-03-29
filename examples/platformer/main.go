package main

import (
	"strings"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/platformer/assets"
	"github.com/dfirebaugh/tortuga/examples/platformer/consumable"
	"github.com/dfirebaugh/tortuga/examples/platformer/heart"
	"github.com/dfirebaugh/tortuga/examples/platformer/player"
	"github.com/dfirebaugh/tortuga/pkg/entity"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

type cart struct{}

var (
	game    tortuga.Console
	player1 entity.Entity
	powerup entity.Entity
)

func (c cart) Update() {
	for _, c := range consumable.Consumables {
		c.Update()
	}
	player1.Update()
	// powerup.Update()
}

func (c cart) Render() {
	game.Clear()
	// game.RenderPalette()
	// powerup.Render()
	for _, c := range consumable.Consumables {
		c.Render()
	}
	player1.Render()
}

func initPlayer(s sprite.Sprite) {
	player1 = player.New(game, s, getPlayerCollidables())
}

func getPlayerCollidables() []geom.Rect {
	c := []geom.Rect{}
	for i, row := range strings.Split(assets.Assets.GetBackgrounds().Get("platforms"), "\n") {
		for j, char := range row {
			if char != '#' && char != '$' {
				continue
			}
			c = append(c, geom.MakeRect(
				float64(j*game.GetTileSize()),
				float64(i*game.GetTileSize()),
				float64(game.GetTileSize()),
				float64(game.GetTileSize())))
		}
	}
	return c
}

func main() {
	game = tortuga.New()

	// if you don't call SetPalette a default palette will be used
	game.SetPalette(assets.Assets.GetPalettes().Get("main"))

	// powerup = heart.New(game, assets.Assets.GetSprites().Get("heart"))

	consumable.Consumables = append(consumable.Consumables, heart.New(game, assets.Assets.GetSprites().Get("heart")))
	// powerup.SetCoordinate()

	// drawing a lot of pixels to the screen is expensive.
	//  It's good to draw the tiles ahead of time and only
	//  update them when necessary
	game.SetTiles(assets.Assets.GetTileDefinitions(), assets.Assets.GetBackgrounds().Get("platforms"))

	initPlayer(assets.Assets.GetSprites().Get("player"))
	game.SetFPSEnabled(true)
	game.SetScaleFactor(3)
	game.Run(cart{})
}
