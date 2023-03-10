# Textures
It's more efficient to render something to a texture and then move that texture around than it is to render directly to the frame buffer on every render call.


## Textures as a display target
`Textures` also implement `display.Displayer` which means that most of tortuga's drawing libraries can draw to a texture.

## Rendering Textures
Textures can be pushed into the render pipeline to be queued for rendering.

> Note that textures are rendered in the order that you push them into the render pipeline.

e.g.
```golang
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
	t      = texture.New(int(circle.Diameter()*2)+1, int(circle.Diameter()*2)+1)
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

```

### Simple Example
```golang
package main

type cart struct {
}

var (
	game tortuga.Console
)

func (c cart) Update() {}

func (c cart) Render() {}

var (
	heartPixels = sprite.Parse("0880088088888888888887888888888888888888088888800088880000088000")
	heartTexture = texture.New(8, 8)
)

func main() {
	game = tortuga.New()

	heartTexture.X = float64(game.GetScreenWidth() / 2)
	heartTexture.Y = float64(game.GetScreenHeight() / 2)

	heartTexture.SetPix(heartPixels)
	game.AddToRenderPipeline(heartTexture)

	game.Run(cart{})
}
```
