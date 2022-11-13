package predator

import (
	"math"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Predator struct {
	component.Coordinate
	component.Width
	Game      tortuga.Console
	Direction float64
	speed     float64
	tick      int
}

const (
	colorIndex = 4
)

func (p *Predator) Update() {
	p.initDir()
	p.initSpeed()
	p.checkBounds()
	p.move()
	p.tick++
	p.checkDeath()
}

func (p *Predator) Render() {
	normal := geom.MakePoint(p.X, p.Y)
	normalLength := p.Width * 4
	v := geom.MakeVector(math.Sin(p.Direction)*float64(normalLength), math.Cos(p.Direction)*float64(normalLength))
	n := normal.ToVector().Add(v)

	geom.MakeLine(
		geom.MakePoint(p.X, p.Y),
		n.ToPoint(),
	).Draw(p.Game.GetDisplay(), p.Game.Color(colorIndex))

	geom.MakeCircle(
		p.X,
		p.Y,
		float64(p.Width)).Filled(p.Game.GetDisplay(), p.Game.Color(colorIndex))
}

func (p *Predator) checkBounds() {
	if p.X-float64(p.Width*2) < 0 || p.X+float64(p.Width*2) > float64(p.Game.GetScreenWidth()) {
		p.Direction += 1.8
		return
	}
	if p.Y-float64(p.Width*2) < 0 || int(p.Y)+int(p.Width*2) > p.Game.GetScreenHeight() {
		p.Direction += 1.8
		return
	}
}

func (p *Predator) initDir() {
	if p.Direction == 0 {
		p.Direction = .2
	}

}
func (p *Predator) initSpeed() {
	if p.speed == 0 {
		p.speed = 1
	}
}
func (p *Predator) move() {
	vx := math.Sin(p.Direction)
	vy := math.Cos(p.Direction)
	vx *= p.speed
	vy *= p.speed

	p.X += vx
	p.Y += vy
}

func (p *Predator) checkDeath() {

}
