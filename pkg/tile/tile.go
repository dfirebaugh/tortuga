package tile

import (
	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/display"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

type Tile struct {
	component.Coordinate
	Pixels    []uint8
	Width     int
	PixelSize float64
}

func Decode(encodedString string) Tile {
	return Tile{
		Pixels:    sprite.Decode(encodedString),
		Width:     8,
		PixelSize: 1,
	}
}

func (t Tile) Draw(dsp display.Displayer) {
	for i, p := range t.Pixels {
		geom.MakeRect(
			t.X+float64((i%t.Width)*int(t.PixelSize)),
			t.Y+float64((i/t.Width)*int(t.PixelSize)),
			t.PixelSize,
			t.PixelSize,
		).
			Filled(dsp, config.Config.Color(p))
	}
}
