package main

import (
	"math/rand"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/prey/entities"
	"github.com/dfirebaugh/tortuga/examples/prey/predator"
	"github.com/dfirebaugh/tortuga/examples/prey/prey"
	"github.com/dfirebaugh/tortuga/pkg/component"
)

type cart struct {
	game tortuga.Console
}

func (c cart) Update() {
	for _, e := range entities.Entities {
		e.Update()
	}
}

func (c cart) Render() {
	c.game.Clear()
	// c.game.RenderPalette()
	for _, e := range entities.Entities {
		e.Render()
	}
}

func main() {
	game := tortuga.New()

	game.SetFPSEnabled(true)

	for i := 0; i < 20; i++ {
		x := float64(rand.Intn(game.GetScreenWidth()) - 20)
		y := float64(rand.Intn(game.GetScreenHeight()) - 20)

		if x < 0 {
			x = float64(game.GetScreenWidth() / 2)
		}

		if y < 0 {
			y = float64(game.GetScreenHeight() / 2)
		}

		entities.Entities = append(entities.Entities, &prey.Prey{
			Game: game,
			Coordinate: component.Coordinate{
				X: x,
				Y: y,
			},
			Width:     component.Width(float64((rand.Intn(300) + 50) / 100)),
			Direction: float64(rand.Intn(300) / 100),
		})
	}

	entities.Entities = append(entities.Entities, &predator.Predator{
		Game: game,
		Coordinate: component.Coordinate{
			X: float64(game.GetScreenWidth()) / 2,
			Y: float64(game.GetScreenHeight()) / 2,
		},
		Width:     .5,
		Direction: float64(rand.Intn(300) / 100),
	})

	game.Run(cart{
		game: game,
	})
}
