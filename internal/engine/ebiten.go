//go:build !gl
// +build !gl

package engine

import "github.com/dfirebaugh/tortuga/internal/engine/ebiten"

func New(emulator ebiten.GameConsole, config ebiten.Config) *ebiten.Game {
	return ebiten.New(emulator, config)
}
