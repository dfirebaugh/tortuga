package ppu

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/internal/emulator/devices/display"
)

type vram interface {
	Swap()
	Put(x int16, y int16, c uint8)
	SetPixel(x, y int16, c color.RGBA)
	Display() error
	Size() (int16, int16)
	Clear()
	GetBuffer() [Width][Height][3]uint8
}

type tilememory interface {
	GetFrame(x, y int) []byte
	SetTile(x, y int, pixels []uint8)
	SetTiles(tileMap map[rune][]uint8, background string)
}

type offset struct {
	X int
	Y int
}

type PPU struct {
	vram         vram
	Transparent  uint8
	Layers       map[Layer]GraphicsLayer
	currentLayer Layer
	tilememory   tilememory
	offsets      []offset
}

func New(v vram, t tilememory) *PPU {
	p := &PPU{
		vram:         v,
		tilememory:   t,
		currentLayer: SpriteLayer,
		Layers:       map[Layer]GraphicsLayer{},
	}

	p.Layers[BackgroundLayer] = newGraphicsLayer()
	p.Layers[SpriteLayer] = newGraphicsLayer()
	p.Layers[WindowLayer] = newGraphicsLayer()
	for range p.Layers {
		p.offsets = append(p.offsets, offset{})
	}
	p.offsets = append(p.offsets, offset{
		X: 10,
		Y: 10,
	})
	return p
}

func (p *PPU) GetFrame() []byte {
	frames := [][]byte{}
	for _, layer := range p.Layers {
		frames = append(frames, layer.GetFrame())
	}
	frames = append(frames, p.tilememory.GetFrame(p.offsets[3].X, p.offsets[3].Y))
	return p.mergeLayers(frames...)
}

func (p *PPU) SetTiles(tileMap map[rune][]uint8, background string) {
	p.tilememory.SetTiles(tileMap, background)
}

func (p *PPU) mergeLayers(frames ...[]uint8) []byte {
	result := frames[0]

	for _, f := range frames {
		for j, pixel := range f {
			if pixel == 0 {
				continue
			}
			result[j] = pixel
		}
	}
	return result
}

func (p *PPU) Put(x, y int16, c uint8) {
	p.vram.Put(x, y, c)
}

func (p *PPU) Swap() {
	p.vram.Swap()
	p.Layers[p.currentLayer] = GraphicsLayer(p.vram.GetBuffer())
}

func (p *PPU) Clear() {
	p.vram.Clear()
}

func (p *PPU) GetDisplay() display.Displayer {
	return p.vram
}

func (p *PPU) ShiftLayer(layer int, x int, y int) {
	p.offsets[layer].X = x
	p.offsets[layer].Y = y
}

func (p *PPU) SetTile(x, y int, pixels []uint8) {
	p.tilememory.SetTile(x, y, pixels)
}
