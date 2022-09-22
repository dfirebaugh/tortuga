package tilememory

import (
	"image/color"
	"strings"
)

const (
	height = 240
	width  = 320
)

type config interface {
	GetTileSize() int
	GetScreenHeight() int
	GetScreenWidth() int
	GetPalette() []color.Color
}

type tilememory struct {
	config config
	buffer [width][height][3]uint8
	frame  []byte
}

func New(c config) *tilememory {
	return &tilememory{
		config: c,
	}
}

func (t *tilememory) SetTiles(tileMap map[rune][]uint8, background string) {
	for rowIndex, row := range strings.Split(background, "\n") {
		for i, char := range row {
			if pixels, ok := tileMap[char]; ok {
				t.SetTile(i, rowIndex, pixels)
			}
		}
	}
}

func (t *tilememory) SetTile(tileX, tileY int, pixels []uint8) {
	tile{
		config: t.config,
		buffer: &t.buffer,
		x:      tileX,
		y:      tileY,
		pixels: pixels,
	}.setPixels()

	t.compileFrame()
}

func (t *tilememory) compileFrame() {
	h := t.config.GetScreenHeight()
	w := t.config.GetScreenWidth()
	var frame []byte = make([]byte, 0, 4*h*w)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			frame = append(frame, t.buffer[x][y][0])
			frame = append(frame, t.buffer[x][y][1])
			frame = append(frame, t.buffer[x][y][2])
			frame = append(frame, 0xFF)
		}
	}
	t.frame = frame
}

func (t tilememory) GetFrame(x int, y int) []byte {
	return t.frame
}
