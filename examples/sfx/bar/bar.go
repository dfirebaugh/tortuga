package bar

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/input"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Bar struct {
	Game tortuga.Console
	component.Coordinate
	component.Width
	Value   float64
	Padding float64
}

func (b *Bar) Update() {
	b.HandleDrag()
}

func (b *Bar) Render() {
	b.Width = 10
	b.Padding = 2

	x := b.X + b.Padding
	y := (float64(b.Game.GetScreenHeight()) / 2) - (b.Value * 100)
	w := float64(b.Width) - b.Padding
	h := b.Value * 100
	geom.MakeRect(x, y, w, h).Filled(b.Game.GetDisplay(), b.Game.Color(3))

	if b.Value != 0 {
		geom.MakeRect(x, y, w, 5).Filled(b.Game.GetDisplay(), b.Game.Color(4))
	}

	geom.MakeLine(geom.MakePoint(0, float64(b.Game.GetScreenHeight())/2), geom.MakePoint(float64(b.Game.GetScreenWidth()), float64(b.Game.GetScreenHeight())/2)).Draw(b.Game.GetDisplay(), b.Game.Color(5))
}

func (b *Bar) IsWithinBounds(coordinate component.Coordinate) bool {
	if coordinate.X <= b.X || coordinate.X >= b.X+float64(int(b.Width)) {
		return false
	}
	if coordinate.Y < 10 {
		return false
	}
	return true
}

func (b *Bar) GetValue() float64 {
	return b.Value
}

func (b *Bar) HandleDrag() {
	if input.IsLeftClickPressed() {
		x, y := input.CursorPositionFloat()
		if !b.IsWithinBounds(component.Coordinate{X: x, Y: y}) {
			return
		}

		a := float64(b.Game.GetScreenHeight() / 2)
		diff := a - y
		b.Value = (diff / a)
		if b.Value < 0 {
			b.Value = 0
		}
	}

	if input.IsRightClickPressed() {
		x, y := input.CursorPositionFloat()
		if !b.IsWithinBounds(component.Coordinate{X: x, Y: y}) {
			return
		}
		b.Value = 0
	}
}
