//go:build webgpu && !ebiten
// +build webgpu,!ebiten

package engine

import (
	"github.com/dfirebaugh/tortuga/internal/engine/game"
	"github.com/dfirebaugh/tortuga/internal/engine/webgpu"
)

func New(emulator game.Console, config game.Config) game.Game {
	return webgpu.New(emulator, config)
}
