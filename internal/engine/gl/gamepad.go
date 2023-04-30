package gl

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyW) == glfw.Press || glfw.GetCurrentContext().GetKey(glfw.KeyUp) == glfw.Press
}
func (Keyboard) IsUpJustPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyW) == glfw.Release || glfw.GetCurrentContext().GetKey(glfw.KeyUp) == glfw.Release
}
func (Keyboard) IsDownPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyS) == glfw.Press || glfw.GetCurrentContext().GetKey(glfw.KeyDown) == glfw.Press
}
func (Keyboard) IsDownJustPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyS) == glfw.Release || glfw.GetCurrentContext().GetKey(glfw.KeyDown) == glfw.Release
}
func (Keyboard) IsLeftPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyA) == glfw.Press || glfw.GetCurrentContext().GetKey(glfw.KeyLeft) == glfw.Press
}
func (Keyboard) IsLeftJustPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyA) == glfw.Release || glfw.GetCurrentContext().GetKey(glfw.KeyLeft) == glfw.Release
}
func (Keyboard) IsRightPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyD) == glfw.Press || glfw.GetCurrentContext().GetKey(glfw.KeyRight) == glfw.Press
}
func (Keyboard) IsRightJustPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyD) == glfw.Release || glfw.GetCurrentContext().GetKey(glfw.KeyRight) == glfw.Release
}
func (Keyboard) IsPrimaryPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyZ) == glfw.Press
}
func (Keyboard) IsSecondaryPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyX) == glfw.Press
}
func (Keyboard) IsPrimaryJustPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyZ) == glfw.Release
}
func (Keyboard) IsSecondaryJustPressed() bool {
	return glfw.GetCurrentContext().GetKey(glfw.KeyX) == glfw.Release
}

type Mouse struct{}

func (m Mouse) CursorPosition() (int, int) {
	x, y := glfw.GetCurrentContext().GetCursorPos()
	return int(x), int(y)
}

func (m Mouse) CursorPositionFloat() (float64, float64) {
	return glfw.GetCurrentContext().GetCursorPos()
}

func (m Mouse) IsLeftClickPressed() bool {
	return glfw.GetCurrentContext().GetMouseButton(glfw.MouseButton1) == glfw.Press
}
func (m Mouse) IsRightClickPressed() bool {
	return glfw.GetCurrentContext().GetMouseButton(glfw.MouseButton2) == glfw.Press
}
func (m Mouse) IsLeftClickJustPressed() bool {
	return glfw.GetCurrentContext().GetMouseButton(glfw.MouseButton1) == glfw.Release
}
func (m Mouse) IsRightClickJustPressed() bool {
	return glfw.GetCurrentContext().GetMouseButton(glfw.MouseButton2) == glfw.Release
}
