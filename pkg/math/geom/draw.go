package geom

import "image/color"

type Drawable interface {
	Draw(d displayer, clr color.Color)
	Filled(d displayer, clr color.Color)
}

func Draw[T Drawable](shape T, d displayer, clr color.Color) {
	shape.Draw(d, clr)
}

func Fill[T Drawable](shape T, d displayer, clr color.Color) {
	shape.Filled(d, clr)
}
