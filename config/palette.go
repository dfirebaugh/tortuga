package config

import (
	"image/color"
)

type Color uint8

type Palette []color.Color

func NewPalette() Palette {
	return Config.Palette
}

func (p Palette) GetColor(i uint8) color.Color {
	return p[i%uint8(len(p))]
}

func (p Palette) RGBA(i uint8) color.RGBA {
	r, g, b, _ := p.GetColor(i).RGBA()

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}
}
