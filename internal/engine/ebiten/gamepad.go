package ebiten

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)
}
func (Keyboard) IsUpJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp)
}
func (Keyboard) IsDownPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)
}
func (Keyboard) IsDownJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown)
}
func (Keyboard) IsLeftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
}
func (Keyboard) IsLeftJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)
}
func (Keyboard) IsRightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)
}
func (Keyboard) IsRightJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)
}
func (Keyboard) IsPrimaryPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyZ)
}
func (Keyboard) IsSecondaryPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyX)
}
func (Keyboard) IsPrimaryJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyZ)
}
func (Keyboard) IsSecondaryJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyX)
}

type Mouse struct{}

func (m Mouse) CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func (m Mouse) CursorPositionFloat() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x), float64(y)
}

func (m Mouse) IsLeftClickPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}
func (m Mouse) IsRightClickPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}
func (m Mouse) IsLeftClickJustPressed() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}
func (m Mouse) IsRightClickJustPressed() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
}
