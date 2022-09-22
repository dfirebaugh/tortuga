package consumable

import (
	"tortuga/pkg/math/geom"
)

type Consumable interface {
	Update()
	Render()
	Consume() string
	HitBox() geom.Rect
}

const (
	ResetJumpCount  = "reset jumpcount"
	AlreadyConsumed = "already consumed"
)

var Consumables = []Consumable{}
