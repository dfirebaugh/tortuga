# Map
In the [wad file](./wad.md), you can define a representation between a character and a tile.

This provides an easy way to push a lot of tiles into tile memory.

In the [wad file](./wad.md), we have this representation configured:

```yaml
tiles:
    '#': examples/platformer/assets/brick.tile
    '$': examples/platformer/assets/block.tile
backgrounds:
    platforms: examples/platformer/assets/background.map
```

We can also use this file to build a collision map.

e.g.
```golang
// build a list of rectangles that somethign could collide against
func getCollidables() []geom.Rect {
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
```

### Example

```
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
$......................................$
$......................................$
$......................................$
$....................###...............$
$......................................$
$................###...................$
$......................................$
$............##############............$
$......................................$
$......................................$
$......################################$
$......................................$
$......................................$
$###...................................$
$......................................$
$.......#########......................$
$..................@...................$
$......................................$
$....................#######...........$
$......................................$
$......................................$
$.............................#####....$
$............######....................$
$......................................$
$....####............#######...........$
$......................................$
$......................................$
$......................................$
$######################################$
```
