package raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	return rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)
}
func (Keyboard) IsUpJustPressed() bool {
	return rl.IsKeyUp(rl.KeyW) || rl.IsKeyUp(rl.KeyUp)
}
func (Keyboard) IsDownPressed() bool {
	return rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown)
}
func (Keyboard) IsDownJustPressed() bool {
	return rl.IsKeyUp(rl.KeyS) || rl.IsKeyUp(rl.KeyDown)
}
func (Keyboard) IsLeftPressed() bool {
	return rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)
}
func (Keyboard) IsLeftJustPressed() bool {
	return rl.IsKeyUp(rl.KeyA) || rl.IsKeyUp(rl.KeyLeft)
}
func (Keyboard) IsRightPressed() bool {
	return rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)
}
func (Keyboard) IsRightJustPressed() bool {
	return rl.IsKeyUp(rl.KeyD) || rl.IsKeyUp(rl.KeyRight)
}
func (Keyboard) IsPrimaryPressed() bool {
	return rl.IsKeyPressed(rl.KeyZ)
}
func (Keyboard) IsSecondaryPressed() bool {
	return rl.IsKeyPressed(rl.KeyX)
}
func (Keyboard) IsPrimaryJustPressed() bool {
	return rl.IsKeyUp(rl.KeyZ)
}
func (Keyboard) IsSecondaryJustPressed() bool {
	return rl.IsKeyUp(rl.KeyX)
}

type Mouse struct{}

func (m Mouse) CursorPosition() (int, int) {
	return int(rl.GetMousePosition().X), int(rl.GetMousePosition().Y)
}

func (m Mouse) CursorPositionFloat() (float64, float64) {
	x, y := m.CursorPosition()
	return float64(x), float64(y)
}

func (m Mouse) IsLeftClickPressed() bool {
	return rl.IsMouseButtonDown(rl.MouseLeftButton)
}
func (m Mouse) IsRightClickPressed() bool {
	return rl.IsMouseButtonDown(rl.MouseRightButton)
}
func (m Mouse) IsLeftClickJustPressed() bool {
	return rl.IsMouseButtonUp(rl.MouseLeftButton)
}
func (m Mouse) IsRightClickJustPressed() bool {
	return rl.IsMouseButtonUp(rl.MouseRightButton)
}
