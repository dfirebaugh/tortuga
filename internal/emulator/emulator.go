package emulator

import (
	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/texture"
)

type Emulator struct {
	cart cart
	clock
	fontProcessingUnit
	dsp
	frameBuffer
	configuration
	renderPipeline renderPipeline
	tileMemory
}

func New(fp fontProcessingUnit, clock clock, dsp dsp, frameBuffer frameBuffer, tileMemory tileMemory, rp renderPipeline) *Emulator {
	return &Emulator{
		fontProcessingUnit: fp,
		clock:              clock,
		dsp:                dsp,
		frameBuffer:        frameBuffer,
		configuration:      config.Config,
		tileMemory:         tileMemory,
		renderPipeline:     rp,
	}
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

func (e Emulator) GetRenderPipeline() []*texture.Texture {
	return e.renderPipeline.Get()
}

func (e Emulator) AddToRenderPipeline(t *texture.Texture) {
	e.renderPipeline.Append(t)
}
func (e Emulator) ClearRenderPipeline() {
	e.renderPipeline.Clear()
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

func (e *Emulator) FillDisplay(c uint8) {
	geom.MakeRect(0, 0, float64(e.GetScreenWidth()), float64(e.GetScreenHeight())).Filled(e.GetDisplay(), e.Color(c))
}
