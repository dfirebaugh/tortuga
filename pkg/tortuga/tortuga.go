package tortuga

import (
	"image/color"
	"tortuga/config"
	"tortuga/internal/emulator"
	"tortuga/internal/emulator/devices/clock"
	"tortuga/internal/emulator/devices/font"
	"tortuga/internal/emulator/devices/ppu"
	"tortuga/internal/emulator/devices/ppu/palette"
	"tortuga/internal/emulator/devices/tilememory"
	"tortuga/internal/emulator/devices/vram"
	"tortuga/internal/engine"
	"tortuga/pkg/math/geom"

	"tinygo.org/x/tinyfont/proggy"
)

type Cart interface {
	Update()
	Render()
}

// Console represents a game system - aka a game console
type Console struct {
	*emulator.Emulator
}

var conf = config.Default()

func (c Console) Run(cart Cart) {
	c.LoadCart(cart)
	engine.New(c, conf).Run()
}

func New() Console {
	c := config.Default()
	p := palette.New(c)
	v := vram.New(c, p)
	return Console{
		emulator.New(
			font.New(v, p, &proggy.TinySZ8pt7b),
			ppu.New(v, tilememory.New(c)),
			clock.New(),
			conf,
		)}
}

func (c Console) RenderPalette() {
	p := c.GetPalette()
	size := float64(c.GetScreenWidth() / len(p))
	for i, _ := range p {
		geom.MakeRect(
			float64(i)*size,
			float64(c.GetScreenHeight())-size, size, size).
			Render(c.GetDisplay(), c.Color(uint8(i)))
	}
}

func (c Console) Color(clr uint8) color.Color {
	return c.GetPalette()[clr]
}
