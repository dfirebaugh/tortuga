package engine

import (
	"github.com/dfirebaugh/tortuga/internal/engine/ebiten"
	"github.com/dfirebaugh/tortuga/internal/engine/game"
)

func New(emulator game.Console) game.Game {
	return ebiten.New(emulator)
}

func GamePad() game.GamePad {
	return ebiten.Keyboard{}
}

func Mouse() game.Mouse {
	return ebiten.Mouse{}
}
