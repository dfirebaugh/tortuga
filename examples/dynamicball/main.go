package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/dynamicball/ball"
	"github.com/dfirebaugh/tortuga/pkg/input"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct {
}

var (
	game  tortuga.Console
	balls []*ball.Ball
)

func (c cart) Update() {
	if input.IsLeftClickPressed() {
		x, y := input.CursorPositionFloat()
		generateBallsAt(1, x, y)
	}

	if input.IsRightClickPressed() {
		balls = []*ball.Ball{}
	}
	for _, ball := range balls {
		ball.Update()
	}
}

func (c cart) Render() {
	game.Clear()
	for _, ball := range balls {
		ball.Render()
	}

}

func generateBallsAt(n int, x float64, y float64) {
	for i := 0; i < n; i++ {
		b := ball.New(game, balls)
		b.X = x
		b.Y = y
		for _, existing := range balls {
			if existing.ContainsPoint(geom.MakePoint(b.X, b.Y)) {
				return
			}
		}

		balls = append(balls, b)
	}
}

func main() {
	game = tortuga.New()
	generateBallsAt(2, float64(game.GetScreenWidth()/2), float64(game.GetScreenHeight()/2))
	game.SetFPSEnabled(true)

	game.Run(cart{})
}
