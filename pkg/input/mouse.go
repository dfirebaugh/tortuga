package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func IsLeftClickPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}
func IsRightClickPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}
