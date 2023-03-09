package main

import (
	"github.com/dfirebaugh/tortuga"
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
var heartTexture = texture.New(texture.Rect(0, 0, 20, 20))

func (c cart) Render() {
}

func main() {
	game = tortuga.New()
	game.SetFPSEnabled(true)
	game.SetRenderPipelineDebug(true)
	game.SetScaleFactor(3)
	heartTexture.X = float64(game.GetScreenWidth() / 2)
	heartTexture.Y = float64(game.GetScreenHeight() / 2)
	heartTexture.Alpha = 0xFF

	sprite.DrawPixels(heartTexture, s, 0, 0)
	heartTexture.Render()
	game.AddToRenderPipeline(heartTexture)

	game.RenderPalette()

	game.Run(cart{})
}
