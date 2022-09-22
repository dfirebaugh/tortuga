package vram

import (
	"image/color"
)

const (
	height = 240
	width  = 320
)

type graphicBuffer [width][height][3]uint8
type bufferLabel uint8

type palette interface {
	GetColor(i uint8) color.Color
}

type config interface {
	GetScreenHeight() int
	GetScreenWidth() int
}

const (
	tmpBuffer bufferLabel = iota
	activeBuffer
)

type VRAM struct {
	buffers []graphicBuffer
	palette palette
	config  config
}

func New(c config, p palette) *VRAM {
	tb := graphicBuffer{}
	ab := graphicBuffer{}
	return &VRAM{
		buffers: []graphicBuffer{tb, ab},
		palette: p,
		config:  c,
	}
}

func (v *VRAM) GetFrame() []byte {
	var frame []byte = make([]byte, 0, 4*height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			frame = append(frame, v.buffers[tmpBuffer][x][y][0])
			frame = append(frame, v.buffers[tmpBuffer][x][y][1])
			frame = append(frame, v.buffers[tmpBuffer][x][y][2])
			frame = append(frame, 0xFF)
		}
	}

	return frame
}

func (v *VRAM) Fill(c uint8) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			r, g, b, _ := v.palette.GetColor(c).RGBA()
			v.buffers[tmpBuffer][col][row][0] = uint8(r) // red
			v.buffers[tmpBuffer][col][row][1] = uint8(g) // green
			v.buffers[tmpBuffer][col][row][2] = uint8(b) // blue
		}
	}
	v.Swap()
}

func (v *VRAM) Clear() {
	v.buffers[tmpBuffer] = graphicBuffer{}
	v.buffers[activeBuffer] = v.buffers[tmpBuffer]
	v.buffers[tmpBuffer] = graphicBuffer{}
}

func (v *VRAM) Swap() {
	v.buffers[activeBuffer] = v.buffers[tmpBuffer]
}

func (v *VRAM) GetBuffer() [width][height][3]uint8 {
	return v.buffers[activeBuffer]
}

func (v *VRAM) Size() (int16, int16) {
	return int16(width), int16(height)
}

func (v *VRAM) Put(x, y int16, color uint8) {
	if x <= 0 || x >= int16(width) || y <= 0 || y >= int16(height) {
		return
	}
	r, g, b, _ := v.palette.GetColor(color).RGBA()
	v.buffers[tmpBuffer][x][y][0] = uint8(r)
	v.buffers[tmpBuffer][x][y][1] = uint8(g)
	v.buffers[tmpBuffer][x][y][2] = uint8(b)
}

func (v *VRAM) SetPixel(x, y int16, c color.RGBA) {
	if x <= 0 || x >= int16(width) || y <= 0 || y >= int16(height) {
		return
	}
	r, g, b, _ := c.RGBA()
	v.buffers[tmpBuffer][x][y][0] = uint8(r)
	v.buffers[tmpBuffer][x][y][1] = uint8(g)
	v.buffers[tmpBuffer][x][y][2] = uint8(b)
}

func (v *VRAM) Display() error {
	return nil
}
