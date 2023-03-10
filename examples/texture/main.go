package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
	"github.com/dfirebaugh/tortuga/pkg/texture"
)

type cart struct {
}

var (
	game tortuga.Console
)

func (c cart) Update() {
}

var s = sprite.Parse("0880088088888888888887888888888888888888088888800088880000088000")
var heartTexture = texture.New(texture.Rect(0, 0, 8, 8))

func (c cart) Render() {
	// setting the background to a different color
	geom.MakeRect(0, 0, float64(game.GetScreenWidth()), float64(game.GetScreenHeight())).Filled(game.GetDisplay(), game.Color(2))
}

func main() {
	game = tortuga.New()
	game.SetFPSEnabled(true)
	game.SetRenderPipelineDebug(true)
	game.SetScaleFactor(3)
	heartTexture.X = float64(game.GetScreenWidth() / 2)
	heartTexture.Y = float64(game.GetScreenHeight() / 2)

	heartTexture.SetPix(s)
	game.AddToRenderPipeline(heartTexture)

	game.SetTile(10, 20, s)
	game.SetTile(10, 10, s)

	game.RenderPalette()

	game.Run(cart{})
}
