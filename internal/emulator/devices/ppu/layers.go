package ppu

const (
	Height                = 240
	Width                 = 320
	BackgroundLayer Layer = iota
	SpriteLayer
	WindowLayer
)

type GraphicsLayer [Width][Height][3]uint8
type Layer uint

func newGraphicsLayer() GraphicsLayer {
	return [Width][Height][3]uint8{}
}

func (gl GraphicsLayer) GetFrame() []byte {
	var frame []byte = make([]byte, 0, 4*Height*Width)
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			frame = append(frame, gl[x][y][0])
			frame = append(frame, gl[x][y][1])
			frame = append(frame, gl[x][y][2])
			frame = append(frame, 0xFF)
		}
	}

	return frame
}
