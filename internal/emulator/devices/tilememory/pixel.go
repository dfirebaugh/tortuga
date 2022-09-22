package tilememory

type pixel struct {
	config config
	buffer *[width][height][3]uint8
	x      int
	y      int
	color  uint8
}

func (p pixel) toBuffer() {
	palette := p.config.GetPalette()
	if p.x >= p.config.GetScreenWidth() || p.y >= p.config.GetScreenHeight() || p.y < 0 || p.x < 0 {
		return
	}
	r, g, b, _ := palette[p.color%uint8(len(palette))].RGBA()
	p.buffer[int16(p.x)][int16(p.y)][0] = uint8(r)
	p.buffer[int16(p.x)][int16(p.y)][1] = uint8(g)
	p.buffer[int16(p.x)][int16(p.y)][2] = uint8(b)
}

func (p pixel) toWorldCoord(tileX, tileY int) (int, int) {
	return (tileX * p.config.GetTileSize()) + p.x, (tileY * p.config.GetTileSize()) + p.y
}
