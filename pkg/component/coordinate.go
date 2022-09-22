package component

type Coordinate struct {
	// x and y represents coordinates on the screen
	X float64
	Y float64
}

func (c *Coordinate) SetCoordinate(newCoord Coordinate) {
	c.X = newCoord.X
	c.Y = newCoord.Y
}
