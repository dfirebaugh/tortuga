package engine

import (
	"github.com/dfirebaugh/tortuga/pkg/engine/ebiten"
	"github.com/dfirebaugh/tortuga/pkg/engine/game"
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
