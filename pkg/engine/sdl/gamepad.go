package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_W)] != 0 || state[uint8(sdl.SCANCODE_UP)] != 0
}

func (Keyboard) IsDownPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_S)] != 0 || state[uint8(sdl.SCANCODE_DOWN)] != 0
}

func (Keyboard) IsLeftPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_A)] != 0 || state[uint8(sdl.SCANCODE_LEFT)] != 0
}

func (Keyboard) IsRightPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_D)] != 0 || state[uint8(sdl.SCANCODE_RIGHT)] != 0
}

func (Keyboard) IsPrimaryPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_Z)] != 0
}

func (Keyboard) IsSecondaryPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_X)] != 0
}

func (Keyboard) IsUpJustPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_W)] != 0 || state[uint8(sdl.SCANCODE_UP)] != 0
}

func (Keyboard) IsDownJustPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_S)] != 0 || state[uint8(sdl.SCANCODE_DOWN)] != 0
}

func (Keyboard) IsLeftJustPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_A)] != 0 || state[uint8(sdl.SCANCODE_LEFT)] != 0
}

func (Keyboard) IsRightJustPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_D)] != 0 || state[uint8(sdl.SCANCODE_RIGHT)] != 0
}

func (Keyboard) IsPrimaryJustPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_Z)] != 0
}

func (Keyboard) IsSecondaryJustPressed() bool {
	state := sdl.GetKeyboardState()
	return state[uint8(sdl.SCANCODE_X)] != 0
}

type Mouse struct{}

func (m Mouse) CursorPosition() (int, int) {
	x, y, _ := sdl.GetMouseState()
	return int(x), int(y)
}

func (m Mouse) CursorPositionFloat() (float64, float64) {
	x, y := m.CursorPosition()
	return float64(x), float64(y)
}

func (m Mouse) IsLeftClickPressed() bool {
	state, _, _ := sdl.GetMouseState()
	return state&sdl.BUTTON_LEFT != 0
}

func (m Mouse) IsRightClickPressed() bool {
	state, _, _ := sdl.GetMouseState()
	return state&sdl.BUTTON_RIGHT != 0
}

func (m Mouse) IsLeftClickJustPressed() bool {
	state, _, _ := sdl.GetMouseState()
	return state&sdl.BUTTON_LEFT == 0
}

func (m Mouse) IsRightClickJustPressed() bool {
	state, _, _ := sdl.GetMouseState()
	return state&sdl.BUTTON_RIGHT == 0
}
