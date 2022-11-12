package main

import (
	"time"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/input"
)

type cart struct {
	input input.PlayerInput
	game  tortuga.Console
}

var (
	happy = []float32{
		264,
		264,
		297,
		264,
		352,
		330,
		264,
		264,
		297,
		264,
		396,
		352,
		264,
		264,
		264,
		440,
		352,
		352,
		330,
		297,
		466,
		466,
		440,
		352,
		396,
		352,
	}

	ode = []string{
		"e4",
		"e4",
		"f4",
		"g4",
		"g4",
		"f4",
		"e4",
		"d4",
		"c4",
		"c4",
		"d4",
		"e4",
		"e4",
		"d4",
		"d4",
		"e4",
		"e4",
		"f4",
		"g4",
		"g4",
		"f4",
		"e4",
		"d4",
		"c4",
		"c4",
		"d4",
		"e4",
		"d4",
		"c4",
		"c4",
	}
)

func (c cart) Update() {
	if c.input.IsDownJustPressed() {
		go func() {
			c.game.PlayNotes(ode, time.Millisecond*550)
		}()
	}
	if c.input.IsUpJustPressed() {
		go func() {
			c.game.PlaySequence(happy, time.Millisecond*250)
		}()
	}
}

func (c cart) Render() {
	c.game.PrintAt("press up arrow to play sequence 1", 10, 170, 5)
	c.game.PrintAt("press down arrow to play sequence 2", 10, 180, 5)
}

func main() {
	game := tortuga.New()
	game.Run(cart{
		input: input.Keyboard{},
		game:  game,
	})
}
