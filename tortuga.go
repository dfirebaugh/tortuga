package tortuga

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/internal/emulator"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/clock"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/dsp"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/font"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/ppu"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/ppu/palette"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/tilememory"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/vram"
	"github.com/dfirebaugh/tortuga/internal/engine"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
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
			&dsp.DSP{},
		)}
}

func (c Console) RenderPalette() {
	p := c.GetPalette()
	size := float64(c.GetScreenWidth() / len(p))
	for i, _ := range p {
		geom.MakeRect(
			float64(i)*size,
			float64(c.GetScreenHeight())-size, size, size).
			Filled(c.GetDisplay(), c.Color(uint8(i)))
	}
}

func (c Console) Color(clr uint8) color.Color {
	return c.GetPalette()[clr]
}
func (c Console) RGBA(clr uint8) color.RGBA {
	r, g, b, _ := c.GetPalette()[clr].RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}
}
