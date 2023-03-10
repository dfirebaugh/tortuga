package tortuga

import (
	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/internal/emulator"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/clock"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/dsp"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/font"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/renderpipeline"
	"github.com/dfirebaugh/tortuga/internal/emulator/devices/tilememory"
	"github.com/dfirebaugh/tortuga/internal/engine"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/texture"
	"tinygo.org/x/tinyfont/proggy"
)

type Texture struct {
	texture.Texture
}

type Cart interface {
	Update()
	Render()
}

// Console represents a game system - aka a game console
type Console struct {
	*emulator.Emulator
}

func (c Console) Run(cart Cart) {
	c.LoadCart(cart)
	engine.New(c).Run()
}

func New() Console {
	p := config.NewPalette()
	display := texture.New(texture.Rect(0, 0, config.Config.GetScreenWidth(), config.Config.GetScreenHeight()))
	rp := &renderpipeline.RenderPipeline{}
	console := Console{
		emulator.New(
			font.New(display, p, &proggy.TinySZ8pt7b),
			clock.New(),
			&dsp.DSP{},
			display,
			tilememory.TileMemory{RenderPipeline: rp},
			rp,
		)}
	return console
}

func (c Console) RenderPalette() {
	p := config.Config.GetPalette()
	size := float64(config.Config.GetScreenWidth() / len(p))
	for i := range p {
		geom.MakeRect(
			float64(i)*size,
			float64(config.Config.GetScreenHeight())-size, size, size).
			Filled(c.GetDisplay(), config.NewPalette().GetColor(uint8(i)))
	}
}
