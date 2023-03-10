package main

import (
	"image/color"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/texture"
)

type cart struct {
}

var (
	game tortuga.Console
)

func (c cart) Update() {}

func (c cart) Render() {
	game.FillDisplay(2)
}

var (
	circle = geom.MakeCircle(8, 8, 4)
	t      = texture.New(texture.Rect(0, 0, int(circle.Diameter()*2)+1, int(circle.Diameter()*2)+1))
)

func main() {
	game = tortuga.New()
	game.SetTransparentColor(color.Black)

	t.X = float64(game.GetScreenWidth() / 2)
	t.Y = float64(game.GetScreenHeight() / 2)
	t.Alpha = 0xFF

	circle.Filled(t, game.Color(3))

	game.AddToRenderPipeline(t)

	game.Run(cart{})
}
