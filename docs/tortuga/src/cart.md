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

The most simple implementation of this would look like the following:

```go
package main

import (
	"tortuga/pkg/tortuga"
	"tortuga/pkg/math/geom"
)


type cart struct{}

var (
    game tortuga.Console
    rect geom.Rect
)

func (c cart) Update() {

}

func (c cart) Render() {
    // render a rectangle on the given display as a certain color
    // draw calls need to happen in the render loop
    rect.Render(game.GetDisplay(), 2)
}

func main() {
    // create a rectangle when the app starts (so we don't create on every render loop)
    rect = geom.MakeRect(20, 20, 20, 20)

    // instantiate the game console
    game = tortuga.New()

    // run the cart (note: this is a blocking operation)
    game.Run(cart{})
}
```
