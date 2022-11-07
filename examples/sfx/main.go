package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/sfx/wave"
	"github.com/dfirebaugh/tortuga/pkg/entity"
)

type cart struct {
	Game     tortuga.Console
	entities []entity.Entity
}

func (c cart) Update() {
	for _, e := range c.entities {
		e.Update()
	}
}
func (c cart) Render() {
	c.Game.Clear()
	for _, e := range c.entities {
		e.Render()
	}

	c.Game.PrintAt("1. draw a pattern above", 10, 180, 5)
	c.Game.PrintAt("2. press down arrow to play sequence", 10, 200, 5)
}

func main() {
	game := tortuga.New()

	entities := []entity.Entity{
		wave.New(game),
	}

	game.Run(cart{
		Game:     game,
		entities: entities,
	})
}
