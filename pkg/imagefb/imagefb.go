package imagefb

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/display"
)

type graphicsBuffer [][3]uint8

type ImageFB struct {
	buffer graphicsBuffer
	Width  int
	Height int
	Alpha  uint8
}

func New(width int, height int) *ImageFB {
	return &ImageFB{
		buffer: make(graphicsBuffer, height*width),
		Width:  width,
		Height: height,
		Alpha:  0xFF,
	}
}

func (i *ImageFB) fillTransparent(frame []byte) []byte {
	r, g, b, a := color.Transparent.RGBA()
	frame = append(frame, uint8(r))
	frame = append(frame, uint8(g))
	frame = append(frame, uint8(b))
	frame = append(frame, uint8(a))
	return frame
}

func (i *ImageFB) Put(x, y int16, c uint8) {
	i.SetPixel(x, y, config.NewPalette().RGBA(c))
}

func (i *ImageFB) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || x > int16(i.Width) || y < 0 || y > int16(i.Height) {
		return
	}
	r, g, b, _ := c.RGBA()
	if len(i.buffer) <= i.Width*int(y)+int(x) {
		return
	}
	i.buffer[i.Width*int(y)+int(x)][0] = uint8(r)
	i.buffer[i.Width*int(y)+int(x)][1] = uint8(g)
	i.buffer[i.Width*int(y)+int(x)][2] = uint8(b)
}

func (i ImageFB) Display() error {
	return nil
}

func (i ImageFB) Size() (int16, int16) {
	return int16(i.Height), int16(i.Width)
}

func (i *ImageFB) Clear() {
	i.buffer = make(graphicsBuffer, i.Height*i.Width)
}

func (i *ImageFB) GetFrame() []byte {
	var frame []byte = make([]byte, 0, 4*i.Height*i.Width)
	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			if i.Width*y+x >= len(i.buffer) {
				continue
			}
			el := i.buffer[i.Width*y+x]

			r, g, b, _ := config.Config.GetTransparentColor().RGBA()
			if el[0] == uint8(r) && el[1] == uint8(g) && el[2] == uint8(b) {
				frame = i.fillTransparent(frame)
				continue
			}
			frame = append(frame, el[0])
			frame = append(frame, el[1])
			frame = append(frame, el[2])
			frame = append(frame, i.Alpha)
		}
	}

	return frame
}

func (i *ImageFB) GetDisplay() display.Displayer {
	return i
}
