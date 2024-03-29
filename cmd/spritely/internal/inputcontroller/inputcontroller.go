package inputcontroller

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/widget"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/message"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputController struct {
	Game       tortuga.Console
	Widgets    []widget.Widget
	MessageBus message.MessageBus
}

func (i InputController) Render() {}

func (i InputController) Update() {
	if i.Game.IsLeftClickPressed() {
		x, y := i.Game.CursorPositionFloat()
		coord := component.Coordinate{
			X: x,
			Y: y,
		}
		for _, w := range i.Widgets {
			if !w.IsWithinBounds(coord) {
				continue
			}

			w.SelectElement(coord)
		}
	}
	if i.Game.IsRightClickPressed() {
		x, y := i.Game.CursorPositionFloat()
		coord := component.Coordinate{
			X: x,
			Y: y,
		}
		for _, w := range i.Widgets {
			if !w.IsWithinBounds(coord) {
				continue
			}

			w.AlternateSelectElement(coord)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if i.MessageBus == nil {
			return
		}
		i.MessageBus.Publish(message.Message{
			Topic: topic.SAVE,
		})
		return
	}
}
