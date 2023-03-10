package emulator

import (
	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/texture"
)

type Emulator struct {
	cart cart
	clock
	fontProcessingUnit
	dsp
	frameBuffer
	configuration
	renderPipeline []*texture.Texture
	deprecated
}

func New(fp fontProcessingUnit, clock clock, dsp dsp, frameBuffer frameBuffer) *Emulator {
	e := &Emulator{
		fontProcessingUnit: fp,
		clock:              clock,
		dsp:                dsp,
		frameBuffer:        frameBuffer,
		configuration:      config.Config,
	}

	return e
}

func (e *Emulator) LoadCart(c cart) {
	e.cart = c
}

func (e Emulator) Update() {
	e.cart.Update()
}

func (e Emulator) Render() {
	e.cart.Render()
	e.frameBuffer.Render()
}

// AddToRenderPipeline allows you to push images into a queue that will be rendered
//
//	This is more efficient than rendering directly to the frame buffer
func (e *Emulator) AddToRenderPipeline(img *texture.Texture) {
	e.renderPipeline = append(e.renderPipeline, img)
}

func (e Emulator) GetRenderPipeline() []*texture.Texture {
	return e.renderPipeline
}

func (e *Emulator) ClearRenderPipeline() {
	e.renderPipeline = make([]*texture.Texture, 0)
}

func (e *Emulator) SetScreenHeight(v int) {
	config.Config.SetScreenHeight(v)
	e.ResetFB()
}
func (e *Emulator) SetScreenWidth(v int) {
	config.Config.SetScreenWidth(v)
	e.ResetFB()
}
func (e *Emulator) ResetFB() {
	e.frameBuffer = texture.New(texture.Rect(0, 0, config.Config.GetScreenWidth(), config.Config.GetScreenHeight()))
	e.fontProcessingUnit.ResetDisplay(e.frameBuffer)
}
