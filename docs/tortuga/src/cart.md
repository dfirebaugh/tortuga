# Cart
A cart (aka a game cartridge) is a representation of game code and memory.

A cart should implement the `tortuga.Cart` interface.

```go
type Cart interface {
    Update()
    Render()
}
```

> Note: that some methods can only be called in the `Render` function e.g. most calls that deal with drawing things.

## Example

The most simple implementation of this would look like the following:

> you can run this example locally with the following command

```
go run github.com/dfirebaugh/tortuga/examples/simple
```

<wasm-view height=400 width=530 src="simple.wasm"></wasm-view>

> note: click the canvas to control the rect with WASD

```go
package main

import (
	"github.com/dfirebaugh/tortuga/pkg/input"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga"
)

type cart struct {
	input input.PlayerInput
}

var (
	game  tortuga.Console
	rect  geom.Rect
	speed = 4.0
)

func (c cart) Update() {
	if c.input.IsDownPressed() {
		rect[1] += speed
	}
	if c.input.IsUpPressed() {
		rect[1] -= speed
	}
	if c.input.IsLeftPressed() {
		rect[0] -= speed
	}
	if c.input.IsRightPressed() {
		rect[0] += speed
	}
}

func (c cart) Render() {
	game.Clear()

	// render a rectangle on the given display as a certain color
	// draw calls need to happen in the render loop
	rect.Draw(game.GetDisplay(), game.Color(2))
}

func main() {
	// create a rectangle when the app starts (so we don't create on every render loop)
	rect = geom.MakeRect(20, 20, 20, 20)

	// instantiate the game console
	game = tortuga.New()

	// run the cart (note: this is a blocking operation)
	game.Run(cart{
		input: input.Keyboard{},
	})
}
```
