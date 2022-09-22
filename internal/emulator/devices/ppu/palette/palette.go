package palette

import (
	"image/color"
)

type Color uint8

type config interface {
	GetPalette() []color.Color
}

type palette struct {
	config config
}

func New(c config) palette {
	return palette{
		config: c,
	}
}

func (p palette) GetColor(i uint8) color.Color {
	colors := p.config.GetPalette()
	return colors[i%uint8(len(colors))]
}
