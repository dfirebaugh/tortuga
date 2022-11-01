package palette

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/message"
)

type Palette struct {
	Game tortuga.Console
	component.Coordinate
	MessageBus   message.MessageBus
	Width        int
	ElementSize  int
	CurrentColor uint8
}

// IsWithinBounds will determine if a coordinate exists within the widget.
func (p Palette) IsWithinBounds(coordinate component.Coordinate) bool {
	if coordinate.X <= p.X || coordinate.X >= p.X+float64(p.Width*p.ElementSize) {
		return false
	}
	if coordinate.Y <= p.Y || coordinate.Y >= p.Y+float64(p.Width*p.ElementSize) {
		return false
	}
	return true
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (p Palette) AlternateSelectElement(coordinate component.Coordinate) {
	p.SelectElement(coordinate)
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (p Palette) SelectElement(coordinate component.Coordinate) {
	if !p.IsWithinBounds(coordinate) {
		return
	}

	x := int(coordinate.X-p.X) / p.ElementSize
	y := int(coordinate.Y-p.Y) / p.ElementSize
	if p.MessageBus == nil {
		return
	}
	p.MessageBus.Publish(message.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: uint8(x + p.Width*y),
	})
}

func (p Palette) Render() {
	for i, c := range p.Game.GetPalette() {
		geom.MakeRect(
			p.X+float64((i%p.Width)*p.ElementSize),
			p.Y+float64((i/p.Width)*p.ElementSize),
			float64(p.ElementSize),
			float64(p.ElementSize),
		).
			Filled(p.Game.GetDisplay(), c)

		geom.MakeRect(
			p.X+float64((int(p.CurrentColor)%p.Width)*p.ElementSize),
			p.Y+float64((int(p.CurrentColor)/p.Width)*p.ElementSize),
			float64(p.ElementSize),
			float64(p.ElementSize),
		).Draw(p.Game.GetDisplay(), p.Game.Color(7))
	}
}
func (p Palette) Update() {}

func (p *Palette) Mailbox() {
	if p.MessageBus == nil {
		return
	}

	msg := p.MessageBus.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.SET_CURRENT_COLOR.String():
			if clr, ok := m.GetPayload().(uint8); ok {
				p.CurrentColor = clr
			}
		}
	}
}
