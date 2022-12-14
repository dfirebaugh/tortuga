package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func CursorPositionFloat() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x), float64(y)
}

func IsLeftClickPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}
func IsRightClickPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}
