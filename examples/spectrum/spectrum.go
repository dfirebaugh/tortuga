package main

import (
	"image/color"
	"math/rand"

	"github.com/ByteArena/box2d"
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
)

type Shape interface {
	Render()
	Update()
}

var (
	game tortuga.Console
)

const (
	pixelPerMeter = 2.0
)

func main() {
	game = tortuga.New()

	gravity := box2d.MakeB2Vec2(0, -9.8)
	world := box2d.MakeB2World(gravity)

	world.SetGravity(gravity)

	game.SetFPSEnabled(true)

	orbColors := []color.Color{
		game.Color(8),
		game.Color(12),
		game.Color(11),
	}

	var orbs []Shape

	for i := 0; i < 100; i++ {

		oc := i % 3
		orbs = append(orbs, &Orb{
			PhysicsObject: PhysicsObject{
				Size:      component.Size{Height: 2},
				Friction:  0,
				MaxThrust: 1,
				Thrust:    .5,
				Rotation:  component.Rotation(rand.Intn(360)),
			},
			Color: orbColors[oc],
		})
	}

	game.Run(&cart{
		game,
		&Ship{
			PhysicsObject: PhysicsObject{
				Position: component.Position{
					X: float64(game.GetScreenWidth()/2) / pixelPerMeter,
					Y: float64(game.GetScreenHeight()/2) / pixelPerMeter,
				},
				Size: component.Size{
					Width:  2,
					Height: 2,
				},
				Friction:  .0010,
				MaxThrust: 1,
			},
			Color:      game.Color(3),
			lumenosity: make(map[ShipColor]float32),
		},
		orbs,
		&world,
	})
}
