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

func Rect(minX, minY, maxX, maxY int) image.Rectangle {
	return image.Rect(minX, minY, maxX, maxY)
}

func New(rect image.Rectangle) *Texture {
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
