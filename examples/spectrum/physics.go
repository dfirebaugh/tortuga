package main

import "github.com/dfirebaugh/tortuga/pkg/component"

type PhysicsObject struct {
	component.Position
	component.Size
	component.Rotation
	component.Velocity
	Thrust    float64
	MaxThrust float64
	Friction  float64
}

func (p *PhysicsObject) Update() {
	if p.Thrust > 0 {
		p.Thrust -= p.Friction
	}

	p.VX *= p.Friction
	p.VY *= p.Friction

	if p.Thrust > p.MaxThrust {
		p.Thrust = p.MaxThrust
	} else if p.Thrust < -p.MaxThrust {
		p.Thrust = -p.MaxThrust
	}

	if p.X+p.Thrust*p.Rotation.Cos() < 0 {
		p.X = float64(game.GetScreenWidth() / pixelPerMeter)
	}
	if p.X+p.Thrust*p.Rotation.Cos() > float64(game.GetScreenWidth()/pixelPerMeter) {
		p.X = 0
	}

	if p.Y+p.Thrust*p.Rotation.Sin() < 0 {
		p.Y = float64(game.GetScreenHeight() / pixelPerMeter)
	}
	if p.Y+p.Thrust*p.Rotation.Sin() > float64(game.GetScreenHeight()/pixelPerMeter) {
		p.Y = 0
	}

	// Apply thrust in the direction of the ship's rotation
	p.VX += p.Thrust * p.Rotation.Cos()
	p.VY += p.Thrust * p.Rotation.Sin()

	p.X += p.VX
	p.Y += p.VY
}
