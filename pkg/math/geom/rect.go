package geom

import (
	"image/color"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinydraw"
)

// Rect a float64 slice with 4 elements []float64{x, y, width, height}
type Rect [4]float64

func MakeRect(x, y, width, height float64) Rect {
	return Rect{x, y, width, height}
}

func (r Rect) IsAxisAlignedCollision(other Rect) bool {
	ax := r[0]
	ay := r[1]
	aw := r[2]
	ah := r[3]

	bx := other[0]
	by := other[1]
	bw := other[2]
	bh := other[3]

	return ax < bx+bw &&
		ax+aw > bx &&
		ay < by+bh &&
		ah+ay > by
}

func (r Rect) GetOverlap(other Rect) (float64, float64) {
	return (r[0] + r[2]) - other[0], (r[1] + r[3]) - other[1]
}

// Dimensions returns the total number of dimensions
func (r Rect) Dimensions() int {
	return 4
}

func (r Rect) GetCenter() (float64, float64) {
	return (r[0] + r[0] + r[2]) / 2, (r[1] + r[1] + r[3]) / 2
}

// Dimension returns the value of the i-th dimension
func (r Rect) Dimension(i int) float64 {
	return r[i]
}

func (r Rect) Render(d displayer, clr color.Color) {
	color, ok := clr.(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledRectangle(d, int16(r[0]), int16(r[1]), int16(r[2]), int16(r[3]), color)
}
