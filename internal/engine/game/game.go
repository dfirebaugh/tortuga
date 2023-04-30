package game

import (
	"github.com/dfirebaugh/tortuga/pkg/texture"
)

type Console interface {
	Update()
	Render()
	GetFrame() []byte
	GetRenderPipeline() []*texture.Texture
}

type Game interface {
	Run()
}

type GamePad interface {
	IsUpPressed() bool
	IsUpJustPressed() bool
	IsDownPressed() bool
	IsDownJustPressed() bool
	IsLeftPressed() bool
	IsLeftJustPressed() bool
	IsRightPressed() bool
	IsRightJustPressed() bool
	IsPrimaryPressed() bool
	IsSecondaryPressed() bool
	IsPrimaryJustPressed() bool
	IsSecondaryJustPressed() bool
}

type Mouse interface {
	CursorPosition() (int, int)
	CursorPositionFloat() (float64, float64)
	IsLeftClickPressed() bool
	IsRightClickPressed() bool
	IsLeftClickJustPressed() bool
	IsRightClickJustPressed() bool
}
