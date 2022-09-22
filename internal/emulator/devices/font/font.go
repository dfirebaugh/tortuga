package font

import (
	"image/color"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinyfont"
)

type displayer interface {
	// Size returns the current size of the display.
	Size() (x, y int16)

	// SetPizel modifies the internal buffer.
	SetPixel(x, y int16, c color.RGBA)

	// Display sends the buffer (if any) to the screen.
	Display() error
}

type palette interface {
	GetColor(uint8) color.Color
}

type FontProcessingUnit struct {
	display displayer
	palette palette
	font    tinyfont.Fonter
}

func New(d displayer, p palette, f tinyfont.Fonter) *FontProcessingUnit {
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
