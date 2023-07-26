package tortuga

import (
	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/emulator"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/clock"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/dsp"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/font"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/renderpipeline"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/tilememory"
	"github.com/dfirebaugh/tortuga/pkg/engine"
	"github.com/dfirebaugh/tortuga/pkg/engine/game"
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
	game.GamePad
	game.Mouse
}

func (c Console) Run(cart Cart) {
	c.LoadCart(cart)
	e := engine.New(c)
	e.Run()
}

func New() Console {
	p := config.NewPalette()
	display := texture.New(config.Config.GetScreenWidth(), config.Config.GetScreenHeight())
	rp := &renderpipeline.RenderPipeline{}
	console := Console{
		emulator.New(
			font.New(display, p, &proggy.TinySZ8pt7b),
			clock.New(),
			&dsp.DSP{},
			display,
			tilememory.TileMemory{RenderPipeline: rp},
			rp,
		),
		engine.GamePad(),
		engine.Mouse(),
	}
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
