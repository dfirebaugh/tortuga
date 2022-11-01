package frames

import (
	"time"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/file"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/message"
	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

type Frames struct {
	Game       tortuga.Console
	MessageBus message.MessageBus
	Frames     [][]uint8
	PixelSize  float64
	Width      int
	component.Coordinate
	currentFrame int
	currentColor uint8
	isPlaying    bool
}

func (f *Frames) init() {
	if f.Width == 0 {
		f.Width = 8
	}
	if f.PixelSize == 0 {
		f.PixelSize = 4
	}
	if len(f.Frames) == 0 {
		for i := 0; i < 8; i++ {
			f.Frames = append(f.Frames, sprite.Decode("0000000000000000000000000000000000000000000000000000000000000000"))
		}
	}
}

func (f *Frames) Render() {
	f.init()

	for i, frame := range f.Frames {
		for j, p := range frame {
			geom.MakeRect(
				float64(i*(f.Width*int(f.PixelSize)))+f.X+float64((j%f.Width)*int(f.PixelSize)),
				f.Y+float64((j/f.Width)*int(f.PixelSize)),
				f.PixelSize,
				f.PixelSize,
			).
				Filled(f.Game.GetDisplay(), f.Game.Color(p))
		}
	}

	geom.MakeRect(float64(f.currentFrame*f.Width*int(f.PixelSize))+f.X, f.Y, f.PixelSize*float64(f.Width), f.PixelSize*float64(f.Width)).Draw(f.Game.GetDisplay(), f.Game.Color(7))
}

func (f Frames) Update() {}

// IsWithinBounds will determine if a coordinate exists within the widget.
func (f Frames) IsWithinBounds(coordinate component.Coordinate) bool {
	if coordinate.X <= f.X || coordinate.X >= f.X+float64(f.Width*int(f.PixelSize))*float64(len(f.Frames)) {
		return false
	}
	if coordinate.Y <= f.Y || coordinate.Y >= f.Y+float64(f.Width*int(f.PixelSize)) {
		return false
	}
	return true
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (f *Frames) SelectElement(coordinate component.Coordinate) {
	if !f.IsWithinBounds(coordinate) {
		return
	}

	x := (coordinate.X - f.X) / f.PixelSize
	// y := (coordinate.Y - f.Y) / f.PixelSize

	if f.MessageBus == nil {
		return
	}
	f.MessageBus.Publish(message.Message{
		Topic:   topic.PUSH_PIXELS,
		Payload: sprite.Encode(f.Frames[f.currentFrame]),
	})

	f.currentFrame = int(x) / f.Width
}

// SelectElement will Set the element at the passed in coordinates to the active element.
func (f Frames) AlternateSelectElement(coordinate component.Coordinate) {
	if !f.IsWithinBounds(coordinate) {
		return
	}
}

func (f *Frames) switchFrame(i int) {
	f.currentFrame = i
	f.MessageBus.Publish(message.Message{
		Topic:   topic.PUSH_PIXELS,
		Payload: sprite.Encode(f.Frames[f.currentFrame]),
	})
}

var (
	quit   = make(chan struct{})
	ticker *time.Ticker
)

func (f *Frames) playAnimation() {
	ticker = time.NewTicker(85 * time.Millisecond)
	quit = make(chan struct{})
	for {
		select {
		case <-ticker.C:
			f.switchFrame((f.currentFrame + 1) % len(f.Frames))
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func (f *Frames) Mailbox() {
	if f.MessageBus == nil {
		return
	}

	msg := f.MessageBus.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.SET_CURRENT_COLOR.String():
			if clr, ok := m.GetPayload().(uint8); ok {
				f.currentColor = clr
			}
		case topic.SET_PIXEL.String():
			if i, ok := m.GetPayload().(int); ok {
				if len(f.Frames[f.currentFrame]) < i || i >= len(f.Frames[f.currentFrame]) {
					break
				}
				f.Frames[f.currentFrame][i] = f.currentColor
			}
		case topic.PLAY_ANIMATION.String():
			if f.isPlaying {
				continue
			}
			f.isPlaying = true
			f.switchFrame(0)
			go f.playAnimation()
		case topic.STOP_ANIMATION.String():
			if !f.isPlaying {
				continue
			}
			f.isPlaying = false
			f.switchFrame(0)
			close(quit)

		case topic.SAVE.String():
			file.Save(f.Frames)
		}
	}
}
