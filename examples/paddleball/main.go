package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type entity interface {
	Update()
	Render()
}

type cart struct {
	entities []entity
	Game     tortuga.Console
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
}

const (
	paddleSpeed = 3
	ballSpeed   = 1.1
)

var (
	game = tortuga.New()
	ball = Ball{
		Coordinate: component.Coordinate{
			X: float64(game.GetScreenWidth()) / 2,
			Y: float64(game.GetScreenHeight()) / 2,
		},
		Velocity: component.Velocity{
			VX: ballSpeed,
			VY: ballSpeed,
		},
		Game: game,
	}
	player = paddle{
		Game:       game,
		Coordinate: component.Coordinate{X: 10, Y: 10},
		Height:     30,
		Width:      5,
		Mover: playerMover{
			Game: game,
		},
		BallPosition: &ball.Coordinate,
		Speed:        paddleSpeed,
	}
	enemy = paddle{
		Game:       game,
		Coordinate: component.Coordinate{X: float64(game.GetScreenWidth()) - 20, Y: 10},
		Height:     30,
		Width:      5,
		Mover: enemyMover{
			Game: game,
		},
		BallPosition: &ball.Coordinate,
		Speed:        paddleSpeed,
	}
)

func main() {
	game.Run(&cart{
		Game: game,
		entities: []entity{
			&ball,
			&player,
			&enemy,
		},
	})
}

type Ball struct {
	component.Coordinate
	component.Velocity
	Game tortuga.Console
}

func (b *Ball) Reset() {
	b.Coordinate.X = float64(game.GetScreenWidth()) / 2
	b.Coordinate.Y = float64(game.GetScreenHeight()) / 2

	b.VX = ballSpeed
	b.VY = ballSpeed
}

func (b *Ball) Update() {
	if b.Coordinate.Y < 0 {
		b.VY *= -ballSpeed
	}

	if b.Coordinate.Y > float64(b.Game.GetScreenHeight()) {
		b.VY *= -ballSpeed
	}

	if b.Coordinate.X > float64(b.Game.GetScreenWidth()) || b.Coordinate.X < 0 {
		b.Reset()
	}

	collisionBox := geom.MakeRect(b.X-4, b.Y-4, 4, 4)
	if collisionBox.IsAxisAlignedCollision(geom.MakeRect(enemy.X, enemy.Y, enemy.Width, enemy.Height)) ||
		collisionBox.IsAxisAlignedCollision(geom.MakeRect(player.X, player.Y, player.Width, player.Height)) {
		b.VX *= -ballSpeed
	}
	b.Coordinate.Y += b.VY
	b.Coordinate.X += b.VX
}

func (b Ball) Render() {
	geom.MakeCircle(b.X, b.Y, 1).Filled(b.Game.GetDisplay(), b.Game.Color(4))
}

type mover interface {
	Move(paddlePosition component.Coordinate, ballPosition component.Coordinate, speed float64) (x float64, y float64)
}

type playerMover struct {
	Game tortuga.Console
}
type enemyMover struct {
	Game tortuga.Console
}

func (p playerMover) Move(paddlePosition component.Coordinate, ballPosition component.Coordinate, speed float64) (x float64, y float64) {
	if p.Game.IsDownPressed() {
		return 0, speed
	}
	if p.Game.IsUpPressed() {
		return 0, -speed
	}
	if p.Game.IsLeftPressed() {
		return -speed, 0
	}
	if p.Game.IsRightPressed() {
		return speed, 0
	}
	return 0, 0
}

func (e enemyMover) Move(paddlePosition component.Coordinate, ballPosition component.Coordinate, speed float64) (x float64, y float64) {
	if ballPosition.Y-10 > paddlePosition.Y {
		return 0, speed
	}

	return 0, -speed
}

type paddle struct {
	component.Coordinate
	Mover        mover
	Game         tortuga.Console
	Speed        float64
	BallPosition *component.Coordinate
	Height       float64
	Width        float64
}

func (p *paddle) Update() {
	x, y := p.Mover.Move(p.Coordinate, *p.BallPosition, p.Speed)
	p.Coordinate.X += x
	p.Coordinate.Y += y
}
func (p paddle) Render() {
	geom.MakeRect(p.X, p.Y, p.Width, p.Height).Filled(p.Game.GetDisplay(), p.Game.Color(4))
}
