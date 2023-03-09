package texture

import (
	"image"

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

func (i *Texture) SetPix(buffer []byte) {
	i.Pix = buffer
}

func (i *Texture) Render() {
	i.Pix = i.GetFrame()
}
