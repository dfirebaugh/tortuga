package texture

import (
	"image"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/imagefb"
)

type Texture struct {
	image.RGBA
	component.Coordinate
	imagefb.ImageFB
}

func New(width int, height int) *Texture {
	rect := image.Rect(0, 0, width, height)
	return &Texture{
		RGBA:    *image.NewRGBA(rect),
		ImageFB: *imagefb.New(rect.Dx(), rect.Dy()),
	}
}

func (t *Texture) SetPix(buffer []byte) {
	for i, pixel := range buffer {
		t.SetPixel(
			int16(i%t.Width),
			int16(i/t.Width),
			config.NewPalette().RGBA(pixel))
	}
}

func (i *Texture) Render() {
	i.Pix = i.GetFrame()
}
