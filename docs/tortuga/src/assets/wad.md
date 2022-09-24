# WAD
`wad` files (which stands for [where's all the data](https://doomwiki.org/wiki/WAD) represents a mapping of files that contain data used to build a scene.

`wad` files are written in a yaml format.

Each top level key represents a different behavior for how the asset will be marshalled into the game.
File extensions don't actually matter.

### example

```yaml
palettes:
    main: examples/platformer/assets/main.palette
tiles:
    '#': examples/platformer/assets/brick.tile
    '$': examples/platformer/assets/block.tile
sprites:
    player: examples/platformer/assets/player.spr
    heart: examples/platformer/assets/heart.spr
backgrounds:
    platforms: examples/platformer/assets/background.map
```
