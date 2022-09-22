package emulator

import (
	"image/color"
	"tortuga/internal/emulator/devices/display"
)

type (
	config interface {
		GetTitle() string
		GetTileSize() int
		GetScaleFactor() int
		GetScreenHeight() int
		GetScreenWidth() int
		GetPalette() []color.Color
		GetVolume() int
		GetDebugEnabled() bool
		GetFPSEnabled() bool
		SetTitle(v string)
		SetScaleFactor(v int)
		SetScreenHeight(v int)
		SetScreenWidth(v int)
		SetPalette(v []color.Color)
		SetDebugEnabled(v bool)
		SetFPSEnabled(v bool)
		SetVolume(v int)
	}

	cart interface {
		Update()
		Render()
	}

	clock interface {
		GetTick() uint
	}

	fontProcessingUnit interface {
		PrintAt(s string, x int, y int, c uint8)
	}

	pixelProcessingUnit interface {
		GetFrame() []byte
		Put(x, y int16, c uint8)
		Swap()
		Clear()
		// SetTile sets a block of pixels into a tile address
		SetTile(x, y int, pixels []uint8)
		// SetTiles will push multiple tiles into memory based upon a
		//	 tileMap - a map that correlates a rune with a set of pixels
		//   background - a string representation of the tiles (using the
		//	 rune as a representation of the set of pixels that should go there)
		// e.g.
		//		var tileMap = map[rune][]uint8 {
		//			'#': []uint8{1, 3, 2, 4, 5, 9}
		//		}
		//		var background = `
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		...........................................
		//		.....................#######...............
		//		...........................................
		//		...........................................
		//		..............................#####........
		//		.............######........................
		//		...........................................
		//		.....####............#######...............
		//		...........................................
		//		...........................................
		//		...........................................
		//		########################################...
		//		`
		//		SetTiles(tileMap, background)
		SetTiles(tileMap map[rune][]uint8, background string)
		GetDisplay() display.Displayer
		ShiftLayer(layer int, x int, y int)
		// RenderSprite(sprite []uint8, x, y float64)
	}
)
