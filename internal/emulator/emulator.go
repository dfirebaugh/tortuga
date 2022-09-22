package emulator

import (
	"image"
)

type Emulator struct {
	cart cart
	clock
	fontProcessingUnit
	pixelProcessingUnit
	config
}

func New(fp fontProcessingUnit, ppu pixelProcessingUnit, clock clock, config config) *Emulator {
	return &Emulator{
		fontProcessingUnit:  fp,
		pixelProcessingUnit: ppu,
		clock:               clock,
		config:              config,
	}
}

func (e *Emulator) LoadCart(c cart) {
	e.cart = c
}

func (e Emulator) Update() {
	e.cart.Update()
}

func (e Emulator) Render(screen *image.RGBA) {
	screen.Pix = []uint8{}
	screen.Pix = e.pixelProcessingUnit.GetFrame()
	e.cart.Render()
	e.pixelProcessingUnit.Swap()
}
