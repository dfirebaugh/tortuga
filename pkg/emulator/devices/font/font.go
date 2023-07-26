package font

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/display"
	"golang.org/x/image/colornames"
	"tinygo.org/x/tinyfont"
)

type palette interface {
	GetColor(uint8) color.Color
}

type FontProcessingUnit struct {
	display display.Displayer
	palette palette
	font    tinyfont.Fonter
}

func New(d display.Displayer, p palette, f tinyfont.Fonter) *FontProcessingUnit {
	return &FontProcessingUnit{
		display: d,
		palette: p,
		font:    f,
	}
}

func (fp *FontProcessingUnit) PrintAt(s string, x int, y int, c uint8) {
	color, ok := fp.palette.GetColor(c).(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinyfont.WriteLine(fp.display, fp.font, int16(x), int16(y), s, color)
}

func (fp *FontProcessingUnit) ResetDisplay(display display.Displayer) {
	fp.display = display
}
