package tilememory

import (
	"strings"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/texture"
)

type renderPipeline interface {
	Append(t *texture.Texture)
}

type TileMemory struct {
	RenderPipeline renderPipeline
}

func (tm TileMemory) SetTile(tileX, tileY int, pixels []uint8) {
	t := texture.New(config.Config.GetTileSize(), config.Config.GetTileSize())

	t.X = float64(tileX * config.Config.GetTileSize())
	t.Y = float64(tileY * config.Config.GetTileSize())
	t.Alpha = 0xFF
	t.SetPix(pixels)

	tm.RenderPipeline.Append(t)
}

func (tm TileMemory) SetTiles(tileMap map[rune][]uint8, background string) {
	for rowIndex, row := range strings.Split(background, "\n") {
		for i, char := range row {
			if pixels, ok := tileMap[char]; ok {
				tm.SetTile(i, rowIndex, pixels)
			}
		}
	}
}
