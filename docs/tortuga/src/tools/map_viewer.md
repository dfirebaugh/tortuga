# Map Viewer

There is a simple map viewer that can be used to preview maps.

You will have to provide a [wad file](../assets/wad.md) and a key for which tiles should be rendered.

e.g.

```bash
go run github.com/dfirebaugh/tortuga/cmd/mapviewer@latest -wad ./cmd/mapviewer/examples/game.wad -key tiles
```

The wad file will need relative paths to any assets required.
e.g.

```yaml
tiles:
    '#': ./cmd/mapviewer/examples/brick.tile
    '$': ./cmd/mapviewer/examples/block.tile
backgrounds:
    tiles: ./cmd/mapviewer/examples/background.map
```
