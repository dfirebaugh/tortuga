package ball

import (
	"math"
	"math/rand"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Ball struct {
	component.UUID
	component.Velocity
	mass float64
	geom.Circle
	game        tortuga.Console
	collidables []*Ball
	collisions  []geom.Line
	color       uint8
	friction    float64
}

func New(game tortuga.Console, collidables []*Ball) *Ball {
	b := Ball{
		UUID:        component.NewUUID(),
		game:        game,
		color:       uint8(rand.Intn(16)),
		friction:    0.002,
		collidables: collidables,
	}
	if b.color == 0 {
		b.color++
	}
	b.X = float64(rand.Intn(game.GetScreenWidth()))
	b.Y = float64(rand.Intn(game.GetScreenHeight()))
	b.VX = float64(rand.Intn(2)) + 1
	b.VY = float64(rand.Intn(2)) + 1
	b.R = float64(rand.Intn(4)) + 2
	b.mass = b.R
	return &b
}

func (b *Ball) Update() {
	b.clampVelocity()
	b.diminishVelocity()
	b.checkCollisions()
	b.handleBorderCollision()
	b.X += b.VX
	b.Y += b.VY
}

func (b Ball) Render() {
	b.Draw(b.game.GetDisplay(), b.game.Color(b.color))

	for _, l := range b.collisions {
		l.Draw(b.game.GetDisplay(), b.game.Color(8))
	}
}

func (b *Ball) handleBorderCollision() {
	if b.X+(b.R*2.0) > float64(b.game.GetScreenWidth()) || b.X < 0 {
		b.VX = b.VX * -1
	}
	if b.Y+(b.R*2.0) > float64(b.game.GetScreenHeight()) || b.Y < 0 {
		b.VY = b.VY * -1
	}
}

func (b *Ball) clampVelocity() {
	limit := 4.0
	if b.VX > limit {
		b.VX = limit
	}

	if b.VX < -limit {
		b.VX = -limit
	}
	if b.VY > limit {
		b.VY = limit
	}
	if b.VY < -limit {
		b.VY = -limit
	}
}

func (b *Ball) diminishVelocity() {
	if b.VX < 0 {
		b.VX += b.friction
	}
	if b.VX > 0 {
		b.VX -= b.friction
	}
	if b.VY < 0 {
		b.VY += b.friction
	}
	if b.VY > 0 {
		b.VY -= b.friction
	}

	if b.VY == b.friction {
		b.VY = 0
	}
	if b.VX == b.friction {
		b.VX = 0
	}
}

func (b Ball) isThisBall(other Ball) bool {
	return other.UUID == b.UUID
}

func (b *Ball) checkCollisions() {
	b.collisions = []geom.Line{}
	for _, ball := range b.collidables {
		if b.isThisBall(*ball) {
			continue
		}

		if b.HasOverlap(ball.Circle) {
			b.dynamicResolution(ball)
			b.handleCollision(ball)
		}
	}
}

func (b *Ball) handleCollision(t *Ball) {
	b.collisions = append(b.collisions, geom.MakeLine(geom.MakePoint(b.X, b.Y), geom.MakePoint(t.X, t.Y)))

	distance := math.Sqrt((b.X-t.X)*(b.X-t.X) + (b.Y-t.Y)*(b.Y-t.Y))
	overlap := .5 * (distance - (b.R * 2) - (t.R * 2))

	b.X -= overlap * (b.X - t.X) / distance
	b.Y -= overlap * (b.Y - t.Y) / distance

	t.X += overlap * (b.X - t.X) / distance
	t.Y += overlap * (b.Y - t.Y) / distance
}

func (b *Ball) dynamicResolution(t *Ball) {
	distance := math.Sqrt((b.X-t.X)*(b.X-t.X) + (b.Y-t.Y)*(b.Y-t.Y))

	// normal
	nx := (t.X - b.X) / distance
	ny := (t.Y - b.Y) / distance

	// tangent
	tx := -1 * ny
	ty := nx

	// dot product tangent
	dpTan1 := b.VX*nx + b.VY*ny
	dpTan2 := t.VX*nx + t.VY*ny

	dpNorm1 := b.VX*nx + b.VY*ny
	dpNorm2 := t.VX*nx + t.VY*ny

	efficiency := 1.0
	// conservation of momentum in 1D
	m1 := efficiency * (dpNorm1*(b.mass-t.mass) + 2.0*t.mass*dpNorm2) / (b.mass + t.mass)
	m2 := efficiency * (dpNorm1*(t.mass-b.mass) + 2.0*b.mass*dpNorm1) / (b.mass + t.mass)

	b.VX = tx*dpTan1 + nx*m1
	b.VY = ty*dpTan1 + ny*m1
	t.VX = tx*dpTan2 + nx*m2
	t.VY = ty*dpTan2 + ny*m2
}
