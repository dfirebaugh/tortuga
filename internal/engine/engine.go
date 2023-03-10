//go:build !gl && !webgpu
// +build !gl,!webgpu

package engine

import (
	"github.com/dfirebaugh/tortuga/internal/engine/game"
	"github.com/dfirebaugh/tortuga/internal/engine/graphicsdriver/ebiten"
)

func New(emulator game.Console) game.Game {
	return ebiten.New(emulator)
}
