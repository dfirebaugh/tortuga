package geom

import "image/color"

type displayer interface {
	Put(x int16, y int16, c uint8)
	SetPixel(x, y int16, c color.RGBA)
	Display() error
	Size() (int16, int16)
	Clear()
}
