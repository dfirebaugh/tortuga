package sdl

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/pkg/engine/game"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type SDLEngine struct {
	Mouse
	width, height int32
	console       game.Console
	renderer      *sdl.Renderer
}

func New(console game.Console) SDLEngine {
	r := SDLEngine{
		width:   int32(config.Config.GetScreenWidth()),
		height:  int32(config.Config.GetScreenHeight()),
		console: console,
	}

	window, err := sdl.CreateWindow(
		config.Config.Title,
		0,
		0,
		r.width,
		r.height,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	r.renderer = renderer

	return r
}

func (r SDLEngine) Update() {
	r.console.Update()
}

func (r SDLEngine) Render() {
	r.renderer.SetDrawColor(0x2b, 0x2b, 0x2b, 0xff)
	r.renderer.Clear()

	frame := r.console.GetFrame()
	for i, v := range toPixels(frame) {
		r.renderer.SetDrawColor(v.R, v.G, v.B, v.A)
		r.renderer.DrawPoint(int32(i%int(r.width)), int32(i/int(r.width)))
	}

	r.console.Render()

	r.renderer.Present()
}

func toPixels(pixelBytes []byte) []color.RGBA {
	pixels := make([]color.RGBA, len(pixelBytes)/4)
	for i := 0; i < len(pixelBytes)/4; i++ {
		pixels[i] = color.RGBA{
			R: pixelBytes[i*4],
			G: pixelBytes[i*4+1],
			B: pixelBytes[i*4+2],
			A: pixelBytes[i*4+3],
		}
	}
	return pixels
}

func (r SDLEngine) Run() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err := img.Init(img.INIT_PNG); err != nil {
		panic(err)
	}
	defer img.Quit()

	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	defer r.renderer.Destroy()

	for {
		event := sdl.WaitEventTimeout(1000 / 120)
		if event != nil {
			switch event.GetType() {
			case sdl.QUIT:
				return
			}
		}
		r.Update()
		r.Render()
	}
}
