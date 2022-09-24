package heart

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/examples/platformer/consumable"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

type heart struct {
	sprite.Sprite
	game tortuga.Console
	component.Coordinate
	consumed bool
	count    int
}

func New(game tortuga.Console, sprite sprite.Sprite) *heart {
	h := heart{
		game: game,
	}
	h.Animations = sprite.Animations
	h.X = 150
	h.Y = 150
	return &h
}

func (h *heart) Update() {
	if h.count == 100 {
		h.consumed = false
	}
	if h.count > 100 {
		h.count = 0
	}
}
func (h *heart) Render() {
	if h.consumed {
		h.count++
		return
	}
	h.Draw(h.game.GetDisplay(), 0, h.X, h.Y)
}

func (h *heart) Consume() string {
	if !h.consumed {
		h.consumed = true
		return consumable.ResetJumpCount
	}
	return consumable.AlreadyConsumed

}

func (h *heart) HitBox() geom.Rect {
	return geom.MakeRect(h.X, h.Y, 8, 8)
}
