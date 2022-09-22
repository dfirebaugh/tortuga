package tilememory

type tile struct {
	config config
	buffer *[width][height][3]uint8
	x      int
	y      int
	pixels []uint8
}

func (t tile) setPixels() {
	for i, p := range t.pixels {
		pixel := pixel{
			config: t.config,
			buffer: t.buffer,
			color:  p,
			x:      i % t.config.GetTileSize(),
			y:      i / t.config.GetTileSize(),
		}
		pixel.x, pixel.y = pixel.toWorldCoord(t.x, t.y)
		pixel.toBuffer()
	}
}
