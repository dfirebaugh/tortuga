//go:build gl && !webgpu && !ebiten
// +build gl,!webgpu,!ebiten

package engine

import (
	"github.com/dfirebaugh/tortuga/internal/engine/game"
	"github.com/dfirebaugh/tortuga/internal/engine/gl"
)

func New(emulator game.Console) game.Game {
	return gl.New(emulator)
}
