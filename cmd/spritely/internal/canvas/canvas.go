package canvas

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/message"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

type Canvas struct {
	Game       tortuga.Console
	MessageBus message.MessageBus
	component.Coordinate
	Width        int
	PixelSize    float64
	Pixels       []uint8
	CurrentColor uint8
}

func (c *Canvas) setDefault() {
	if len(c.Pixels) == 0 {
		for i := 0; i < 64; i++ {
			c.Pixels = append(c.Pixels, uint8(0))
		}
	}
}
func (c *Canvas) Render() {
	c.setDefault()
	for i, p := range c.Pixels {
		geom.MakeRect(
			c.X+float64((i%c.Width)*int(c.PixelSize)),
			c.Y+float64((i/c.Width)*int(c.PixelSize)),
			c.PixelSize,
			c.PixelSize,
		).
			Filled(c.Game.GetDisplay(), c.Game.Color(p))
	}
}
func (c Canvas) Update() {}

func (c *Canvas) replacePixels(spriteString string) {
	c.Pixels = sprite.Decode(spriteString)
}

// IsWithinBounds will determine if a coordinate exists within the widget.
func (c Canvas) IsWithinBounds(coordinate component.Coordinate) bool {
	if coordinate.X <= c.X || coordinate.X >= c.X+float64(c.Width*int(c.PixelSize)) {
		return false
	}
	if coordinate.Y <= c.Y || coordinate.Y >= c.Y+float64(c.Width*int(c.PixelSize)) {
		return false
	}
	return true
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (c Canvas) SelectElement(coordinate component.Coordinate) {
	if !c.IsWithinBounds(coordinate) {
		return
	}

	x := (coordinate.X - c.X) / c.PixelSize
	y := (coordinate.Y - c.Y) / c.PixelSize

	if c.MessageBus == nil {
		return
	}

	c.MessageBus.Publish(message.Message{
		Topic:   topic.SET_PIXEL,
		Payload: c.Width*int(y) + int(x),
	})
	println(sprite.Encode(c.Pixels))
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (c Canvas) AlternateSelectElement(coordinate component.Coordinate) {
	if !c.IsWithinBounds(coordinate) {
		return
	}

	x := (coordinate.X - c.X) / c.PixelSize
	y := (coordinate.Y - c.Y) / c.PixelSize

	if c.MessageBus == nil {
		return
	}

	c.MessageBus.Publish(message.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: c.Pixels[c.Width*int(y)+int(x)],
	})
}

func (c *Canvas) Mailbox() {
	if c.MessageBus == nil {
		return
	}

	msg := c.MessageBus.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.SET_CURRENT_COLOR.String():
			if clr, ok := m.GetPayload().(uint8); ok {
				c.CurrentColor = clr
			}
		case topic.SET_PIXEL.String():
			if i, ok := m.GetPayload().(int); ok {
				if len(c.Pixels) < i || i >= len(c.Pixels) {
					break
				}
				c.Pixels[i] = c.CurrentColor
			}
		case topic.PUSH_PIXELS.String():
			c.replacePixels(m.GetPayload().(string))
		}
	}
}
