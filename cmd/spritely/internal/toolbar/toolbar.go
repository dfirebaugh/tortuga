package toolbar

import (
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/topic"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/message"
	"github.com/dfirebaugh/tortuga/pkg/tile"
)

type ToolBar struct {
	Game tortuga.Console
	component.Coordinate
	MessageBus  message.MessageBus
	Width       int
	PixelSize   float64
	isPlaying   bool
	currentTool int
}

func (t *ToolBar) Render() {
	play := tile.Decode("7737777772332227723332277233332772333327723332277233222777377777")
	if t.isPlaying {
		play = tile.Decode("7737777776336667763336677633336776333367763336677633666777377777")
	}

	stop := tile.Decode("7777777772222227728888277288882772888827728888277222222777777777")
	save := tile.Decode("22222222d56666d2d566d6ddd56666ddddddddddd777777dd777777dd222222d")
	pencil := tile.Decode("22222ee22222d1ee2229adae229a9ad229a9a9224a9a9222f4a922220f422222")
	bucket := tile.Decode("22022222202021122020011e220601ee206760ee0677760e2067602e22060222")
	hand := tile.Decode("2222222207070702070707020777776007777760077777600777770200000022")
	// key := tile.Decode("aaaaaa22a2222a22aaaaaa2222a2222222aa222222aaa22222a2222222aaa222")

	tiles := []tile.Tile{}
	tiles = append(tiles, play)
	tiles = append(tiles, stop)
	tiles = append(tiles, pencil)
	tiles = append(tiles, bucket)
	tiles = append(tiles, hand)
	tiles = append(tiles, save)
	// tiles = append(tiles, key)

	for i, tile := range tiles {
		tile.X = float64(t.Width) * t.PixelSize * float64(i)
		tile.Y = 130
		tile.PixelSize = t.PixelSize
		tile.Draw(t.Game.GetDisplay())
	}
	geom.MakeRect(
		float64(t.currentTool*t.Width*int(t.PixelSize))+t.X,
		t.Y,
		t.PixelSize*float64(t.Width),
		t.PixelSize*float64(t.Width),
	).
		Draw(t.Game.GetDisplay(), t.Game.Color(7))
}
func (t ToolBar) Update() {}

// IsWithinBounds will determine if a coordinate exists within the widget.
func (t ToolBar) IsWithinBounds(coordinate component.Coordinate) bool {
	if coordinate.X <= t.X || coordinate.X >= t.X+float64(t.Width*int(t.PixelSize))*float64(6) {
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

	x := (coordinate.X - t.X) / t.PixelSize

	t.currentTool = int(x) / t.Width

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
		}
	}
}
