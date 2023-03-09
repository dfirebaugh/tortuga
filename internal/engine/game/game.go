package game

import "github.com/dfirebaugh/tortuga/pkg/texture"

type Console interface {
	Update()
	Render()
	GetFrame() []byte
	GetRenderPipeline() []*texture.Texture
}

type Game interface {
	Run()
}
