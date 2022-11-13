package prey

import (
	"math"
	"math/rand"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/prey/entities"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Prey struct {
	component.Coordinate
	component.Width
	Game      tortuga.Console
	Direction float64
	speed     float64
	tick      int
}

const (
	colorIndex = 3
)

func (p *Prey) Update() {
	p.initDir()
	p.initSpeed()
	p.checkBounds()
	p.move()
	p.tick++
	p.checkGeneration()
}

func (p *Prey) Render() {
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

func (p *Prey) checkBounds() {
	if p.X-float64(p.Width*2) < 0 || p.X+float64(p.Width*2) > float64(p.Game.GetScreenWidth()) {
		p.Direction += 1.8
		return
	}
	if p.Y-float64(p.Width*2) < 0 || int(p.Y)+int(p.Width*2) > p.Game.GetScreenHeight() {
		p.Direction += 1.8
		return
	}
}

func (p *Prey) initDir() {
	if p.Direction == 0 {
		p.Direction = .2
	}

}
func (p *Prey) initSpeed() {
	if p.speed == 0 {
		p.speed = .2
	}
}

func (p *Prey) move() {
	vx := math.Sin(p.Direction)
	vy := math.Cos(p.Direction)
	vx *= p.speed
	vy *= p.speed

	p.X += vx
	p.Y += vy
}

func (p *Prey) checkGeneration() {
	if p.tick%500 == 0 {
		entities.Entities = append(entities.Entities, &Prey{
			Coordinate: component.Coordinate{X: p.X + float64(p.Width*2), Y: p.Y},
			Width:      component.Width(float64((rand.Intn(300) + 50) / 100)),
			Game:       p.Game,
			Direction:  p.Direction + 1.8,
			speed:      0,
			tick:       0,
		})
	}
}
