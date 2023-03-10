package main

import (
	"fmt"
	"math/rand"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct {
}

var (
	game       tortuga.Console
	cells      []*cell
	generation = 0
	tick       = 0
)

func (c cart) Update() {
	if tick%5 == 0 {
		generation += 1
		for _, cell := range cells {
			cell.iteration()
		}
		for _, cell := range cells {
			cell.isAlive = cell.nextGen
		}
	}
	tick++

}
func (c cart) Render() {
	game.Clear()
	geom.MakeRect(0, 0, float64(game.GetScreenWidth()), float64(game.GetScreenHeight())).
		Filled(game.GetDisplay(), game.Color(0))
	game.PrintAt(fmt.Sprintf("gen: %d", generation), 10, 25, 2)
	for _, cell := range cells {
		cell.Render()
	}

}

type cell struct {
	geom.Rect
	isAlive bool
	component.Coordinate
	nextGen bool
}

func (c *cell) Update() {}

func (c cell) Render() {
	if !c.isAlive {
		return
	}

	c.Filled(game.GetDisplay(), game.Color(5))
}

func (c *cell) iteration() {
	if !c.isAlive {
		c.nextGen = c.shouldReproduce()
		return
	}
	c.nextGen = !c.shouldDie()
}

func (c *cell) shouldDie() bool {
	if c.isUnderpopulated() || c.isOverpopulated() {
		return true
	}

	return false
}

func (c cell) isOverpopulated() bool {
	return c.getAliveNeighborCount() > 3
}
func (c cell) isUnderpopulated() bool {
	return c.getAliveNeighborCount() < 2
}
func (c cell) shouldReproduce() bool {
	return c.getAliveNeighborCount() == 3
}

func (c cell) getAliveNeighborCount() int {
	x := c.X
	y := c.Y
	potential := []component.Coordinate{
		{Y: y, X: x - 1},     //left
		{Y: y, X: x + 1},     //right
		{Y: y - 1, X: x},     //up
		{Y: y + 1, X: x},     //down
		{Y: y - 1, X: x - 1}, //left top diagnal
		{Y: y + 1, X: x + 1}, //right bottom diagnal
		{Y: y - 1, X: x + 1}, // left bottom diagnal
		{Y: y + 1, X: x - 1}, // right top diagnal
	}

	neighbors := []bool{}

	for _, n := range potential {
		if n.X < 0 || n.Y < 0 || n.X > float64(game.GetScreenWidth()) || n.Y > float64(game.GetScreenHeight()) {
			continue
		}

		if game.GetScreenWidth()*int(n.Y)+int(n.X) >= len(cells) {
			continue
		}
		if cells[game.GetScreenWidth()*int(n.Y)+int(n.X)].isAlive {
			neighbors = append(neighbors, true)
		}
	}

	return len(neighbors)
}

func main() {
	game = tortuga.New()
	game.SetScaleFactor(3)
	game.SetFPSEnabled(true)

	for y := 0; y < game.GetScreenHeight(); y++ {
		for x := 0; x < game.GetScreenWidth(); x++ {
			c := &cell{}
			c.X = float64(x)
			c.Y = float64(y)
			c.Rect = geom.MakeRect(c.X, c.Y, 1, 1)
			c.isAlive = rand.Intn(100) < 25
			cells = append(cells, c)
		}
	}

	game.Run(cart{})
}
