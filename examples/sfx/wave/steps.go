package wave

import (
	"time"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/sfx/bar"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/input"
)

type Steps struct {
	game    tortuga.Console
	samples []*bar.Bar
}

func New(game tortuga.Console) *Steps {
	width := 10
	samples := []*bar.Bar{}
	for i := 0; i < game.GetScreenWidth()/width; i++ {
		samples = append(samples, &bar.Bar{
			Game: game,
			Coordinate: component.Coordinate{
				X: float64(width) * float64(i),
			},
		})
	}
	return &Steps{
		game:    game,
		samples: samples,
	}
}

func (s *Steps) Update() {
	for _, b := range s.samples {
		b.Update()
	}

	keyboard := input.Keyboard{}
	if keyboard.IsDownJustPressed() {
		go func(samples []*bar.Bar) {
			sequence := []float32{}
			for _, f := range samples {
				sequence = append(sequence, float32(f.GetValue()*1000))
			}

			s.game.PlaySequence(sequence, time.Microsecond*50000)
		}(s.samples)
	}
}

func (s Steps) Render() {
	for _, b := range s.samples {
		b.Render()
	}
}
