package main

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/input"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct {
}

type player struct {
	component.Velocity
	input input.PlayerInput
	rect  geom.Rect
	speed float64
}

var (
	playerRect = &player{
		rect:  geom.MakeRect(float64(game.GetScreenWidth()/2-20), float64(game.GetScreenHeight()/2-20), 40, 40),
		speed: 3,
		input: input.Keyboard{},
	}
	game      = tortuga.New()
	rects     = []geom.Rect{}
	damping   = .1
	ray       = geom.Ray{}
	collision = &geom.Collision{}
)

func (c cart) Render() {
	d := game.GetDisplay()
	game.Clear()

	for _, r := range rects {
		r.Draw(d, game.Color(4))
	}
	playerRect.render()

	if playerRect.VX == 0 && playerRect.VY == 0 {
		return
	}
	for _, r := range rects {
		if r.HasRayIntersection(ray, collision) {
			playerRect.rect.Filled(d, game.Color(4))
			r.Filled(d, game.Color(4))
			geom.MakeLine(
				ray.Origin.ToPoint(),
				collision.Point.ToPoint(),
			).Draw(d, game.Color(6))
			return
		}
	}

	geom.MakeLine(
		ray.Origin.ToPoint(),
		ray.Direction.ToPoint(),
	).Draw(d, game.Color(3))
}

func (c cart) Update() {
	playerRect.update()
	ray = geom.Ray{
		Origin:    geom.MakeVector(playerRect.rect[0]+(playerRect.rect[2]/2), playerRect.rect[1]+(playerRect.rect[3]/2)),
		Direction: geom.MakeVector(playerRect.VX*50, playerRect.VY*50),
	}
}

func (p *player) render() {
	playerRect.rect.Draw(game.GetDisplay(), game.Color(3))
}
func (p *player) update() {
	p.clampVelocity()
	p.diminishVelocity()
	if p.input.IsDownPressed() {
		playerRect.VY += playerRect.speed
	}
	if p.input.IsUpPressed() {
		playerRect.VY -= playerRect.speed
	}
	if p.input.IsLeftPressed() {
		playerRect.VX -= playerRect.speed
	}
	if p.input.IsRightPressed() {
		playerRect.VX += playerRect.speed
	}
	p.rect[0] += p.VX
	p.rect[1] += p.VY
}

func (p *player) clampVelocity() {
	limit := 4.0
	if p.VX > limit {
		p.VX = limit
	}

	if p.VX < -limit {
		p.VX = -limit
	}
	if p.VY > limit {
		p.VY = limit
	}
	if p.VY < -limit {
		p.VY = -limit
	}
}
func (p *player) diminishVelocity() {
	if p.VX < 0 {
		p.VX += damping
	}
	if p.VX > 0 {
		p.VX -= damping
	}
	if p.VY < 0 {
		p.VY += damping
	}
	if p.VY > 0 {
		p.VY -= damping
	}

	if p.VY == damping {
		p.VY = 0
	}
	if p.VX == damping {
		p.VX = 0
	}
}

func main() {
	for i := 0; i < game.GetScreenWidth()/15; i++ {
		rects = append(rects, geom.MakeRect(
			float64(i*game.GetScreenWidth()/15),
			15,
			15,
			15,
		))
		rects = append(rects, geom.MakeRect(
			float64(i*game.GetScreenWidth()/15),
			float64(game.GetScreenHeight()-20),
			15,
			15,
		))
	}
	game.Run(cart{})
}
