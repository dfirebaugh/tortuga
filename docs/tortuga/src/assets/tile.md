# Tile
A tile is represented by hexadecimal digits. Each digit refers to an index on the palette.

## Tile Size
Tiles are 8x8 by default.
you can configure tiles to be of a different size.  However, tiles can only be square.

```golang
// set tiles to be 16x16
game.SetTileSize(16)
```

## Tile Memory
Tile memory exists as a way to store tiles in a layer that doesn't have to rerender very often.

To add a tile to tile memory:
```golang
package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

type cart struct{}

func (c cart) Update() {}
func (c cart) Render() {}

var game tortuga.Console
func main() {
    game = tortuga.New()

    // pushes a tile to tile memory
    game.SetTile(2, 5, sprite.Parse("bbbbbbbbb4444444444455444444544445444444454444444444444444444444"))
    game.Run(cart{})
}
```

## Creating a tile
You could dynamnically create a tile in code.

e.g.
```golang
package main

import "github.com/dfirebaugh/tortuga"

type cart struct{}

func (c cart) Update() {}
func (c cart) Render() {}

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
    game.SetTile(2, 5, blockFactory(7))

    game.Run(cart{})
}
```

### Example
example of a tile file:

```
bbbbbbbbb4444444444455444444544445444444454444444444444444444444
```
