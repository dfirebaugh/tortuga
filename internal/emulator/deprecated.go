package emulator

type deprecated struct{}

// SetTile is deprecated
func (d deprecated) SetTile(tileX, tileY int, pixels []uint8) {}

// SetTiles is deprecated
func (d deprecated) SetTiles(tileMap map[rune][]uint8, background string) {}

// GetTileFrame is deprecated
func (d deprecated) GetTileFrame(x int, y int) []byte {
	return nil
}
