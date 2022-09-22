package player

import (
	"fmt"
	"tortuga/examples/platformer/consumable"
	"tortuga/pkg/component"
	"tortuga/pkg/input"
	"tortuga/pkg/math/geom"
	"tortuga/pkg/sprite"
	"tortuga/pkg/tortuga"
)

type player struct {
	component.Coordinate
	component.Velocity
	sprite.Sprite
	hitbox      geom.Rect
	speed       float64
	input       input.PlayerInput
	game        tortuga.Console
	collidables []geom.Rect
	jumpCount   int
}

var (
	gravity         = 3
	x       float64 = 0
	y       float64 = 0
)

const (
	damping = .1
)

func New(game tortuga.Console, input input.PlayerInput, s sprite.Sprite, collidables []geom.Rect) *player {
	p := &player{
		input:       input,
		game:        game,
		speed:       .3,
		hitbox:      geom.MakeRect(0, 0, float64(8), float64(16)),
		collidables: collidables,
	}
	p.Animations = s.Animations
	p.X = 42
	p.Y = float64((game.GetScreenHeight() / game.GetTileSize()))
	return p
}

func (p *player) Render() {
	p.hitbox.Draw(p.game.GetDisplay(), p.game.Color(1))
	if p.input.IsLeftPressed() {
		p.DrawPixels(p.game.GetDisplay(), p.GetPixels(0, 0), p.X, p.Y)
		return
	}
	p.Draw(p.game.GetDisplay(), 0, p.X, p.Y)
	p.game.PrintAt(fmt.Sprintf("%f", p.VY), int(p.X), int(p.Y), 1)
	p.game.PrintAt(fmt.Sprintf("%d", p.jumpCount), int(p.X), int(p.Y)+25, 3)
	p.game.PrintAt(fmt.Sprintf("%f : %f", x, y), int(p.X)+25, int(p.Y)-25, 4)
}

func (p *player) Eat(c consumable.Consumable) {
	if c.Consume() == consumable.ResetJumpCount {
		p.jumpCount = 0
	}
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

func (p *player) resetJumpCount() {
	p.jumpCount = 0
}

func (p *player) Update() {
	p.clampVelocity()
	p.diminishVelocity()

	if p.input.IsDownPressed() {
		p.VX += p.speed
	}
	if p.input.IsPrimaryJustPressed() {
		if p.jumpCount < 3 {
			tmpGrav := gravity
			gravity = 0
			p.VY -= 15
			p.jumpCount++
			gravity = tmpGrav
		}
	}
	if p.input.IsLeftPressed() {
		p.VX -= p.speed
	}
	if p.input.IsRightPressed() {
		p.VX += p.speed
	}
	p.hitbox[0] = p.X + p.VX
	p.hitbox[1] = p.Y + p.VY
	if p.getCollisionMap(p.hitbox) {
		// p.VX = 0
		// p.VY = 0
		return
	}

	for _, c := range consumable.Consumables {
		if !c.HitBox().IsAxisAlignedCollision(p.hitbox) {
			continue
		}

		p.Eat(c)
	}

	p.X += p.VX
	p.Y += p.VY
	p.applyGravity()
	p.hitbox[0] = p.X
	p.hitbox[1] = p.Y
}

func (p *player) applyGravity() {
	hb := p.hitbox
	hb[1] += float64(gravity)
	if p.getCollisionMap(hb) {
		p.resetJumpCount()
		p.VY = 0
		return
	}
	if p.input.IsPrimaryJustPressed() {
		return
	}
	p.Y += float64(gravity)
}

func (p *player) getCollisionMap(hitbox geom.Rect) bool {
	for _, r := range p.collidables {
		if r.IsAxisAlignedCollision(hitbox) {
			// p.resolveCollision(r)
			// y = (p.hitbox[1] + p.hitbox[3]) - r[1]
			// p.Y = (p.hitbox[1] + p.hitbox[3]) - r[1]
			return true
		}
		x = 0
		y = 0
	}

	return false
}

func (p *player) resolveCollision(r geom.Rect) {
	x, y := p.hitbox.GetOverlap(r)
	// println(x, y)
	p.X -= x / 2
	p.Y -= y / 2
}
