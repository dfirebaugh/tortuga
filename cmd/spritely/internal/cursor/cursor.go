package cursor

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/message"
)

type Cursor struct {
	Game       tortuga.Console
	MessageBus message.MessageBus
	component.Coordinate
	Color uint8
}

func (c Cursor) Render() {
	x, y := c.Game.CursorPosition()
	geom.MakeRect(float64(x), float64(y), 5, 5).Filled(c.Game.GetDisplay(), c.Game.Color(c.Color))
}
func (c Cursor) Update() {}

func (c *Cursor) Mailbox() {
	if c.MessageBus == nil {
		return
	}

	msg := c.MessageBus.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.SET_CURRENT_COLOR.String():
			if paletteIndex, ok := m.GetPayload().(uint8); ok {
				c.Color = paletteIndex
			}
		}
	}
}
