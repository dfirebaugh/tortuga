package toolbar

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/message"
	"github.com/dfirebaugh/tortuga/pkg/tile"
)

type ToolBar struct {
	Game tortuga.Console
	component.Coordinate
	MessageBus message.MessageBus
	Width      int
	PixelSize  float64
	isPlaying  bool
}

func (t *ToolBar) Render() {
	play := tile.Decode("7737777770330007703330077033330770333307703330077033000777377777")
	if t.isPlaying {
		play = tile.Decode("7737777776336667763336677633336776333367763336677633666777377777")
	}
	play.X = 0
	play.Y = 130
	play.PixelSize = t.PixelSize
	play.Draw(t.Game.GetDisplay())
	stop := tile.Decode("7777777770000007708888077088880770888807708888077000000777777777")
	// 7777777776666667768888677688886776888867768888677666666777777777
	stop.X = float64(t.Width) * t.PixelSize
	stop.Y = 130
	stop.PixelSize = t.PixelSize
	stop.Draw(t.Game.GetDisplay())
}
func (t ToolBar) Update() {}

// IsWithinBounds will determine if a coordinate exists within the widget.
func (t ToolBar) IsWithinBounds(coordinate component.Coordinate) bool {
	if coordinate.X <= t.X || coordinate.X >= t.X+float64(t.Width*int(t.PixelSize))*float64(2) {
		return false
	}
	if coordinate.Y <= t.Y || coordinate.Y >= t.Y+float64(t.Width*int(t.PixelSize)) {
		return false
	}
	return true
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (t *ToolBar) SelectElement(coordinate component.Coordinate) {
	if !t.IsWithinBounds(coordinate) {
		return
	}

	// for _, tool := range t.Tools {
	// 	tool.ClickHandler()
	// }
	x := (coordinate.X - t.X) / t.PixelSize
	// // y := (coordinate.Y - t.Y) / t.PixelSize

	// if t.Broker == nil {
	// 	return
	// }
	// t.Broker.Publish(message.Message{
	// 	Topic:   topic.PUSH_PIXELS,
	// 	Payload: sprite.Encode(t.Tools[t.currentFrame]),
	// })

	switch int(x) / t.Width {
	case 0:
		println("play")
		t.isPlaying = true
		t.MessageBus.Publish(message.Message{
			Topic: topic.PLAY_ANIMATION,
		})
	case 1:
		println("stop")
		t.isPlaying = false
		t.MessageBus.Publish(message.Message{
			Topic: topic.STOP_ANIMATION,
		})
	}

	// t.currentFrame = int(x) / t.Width
	// println(int(x) / t.Width)
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (t ToolBar) AlternateSelectElement(coordinate component.Coordinate) {
	if !t.IsWithinBounds(coordinate) {
		return
	}
}

func (t *ToolBar) Mailbox() {
	if t.MessageBus == nil {
		return
	}

	msg := t.MessageBus.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		// case topic.SET_CURRENT_COLOR.String():
		// 	if clr, ok := m.GetPayload().(uint8); ok {
		// 		t.currentColor = clr
		// 	}
		// case topic.SET_PIXEL.String():
		// 	if i, ok := m.GetPayload().(int); ok {
		// 		if len(t.Frames[t.currentFrame]) < i || i >= len(t.Frames[t.currentFrame]) {
		// 			break
		// 		}
		// 		t.Frames[t.currentFrame][i] = t.currentColor
		// 	}
		}
	}
}
