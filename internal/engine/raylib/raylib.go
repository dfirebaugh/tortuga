package raylib

import (
	"image/color"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/internal/engine/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RayLib struct {
	width, height int
	console       game.Console
}

func New(console game.Console) RayLib {
	r := RayLib{
		width:   config.Config.GetScreenWidth(),
		height:  config.Config.GetScreenHeight(),
		console: console,
	}

	return r
}

func (r RayLib) Update() {
	r.console.Update()
}

func (r RayLib) Render() {
	rl.BeginDrawing()

	// Clear the screen
	// rl.ClearBackground(color.RGBA{R: 0x2b, G: 0x2b, B: 0x2b, A: 0xff})

	for i, v := range toPixels(r.console.GetFrame()) {
		rl.DrawPixel(
			int32(i%config.Config.GetScreenWidth()),
			int32(i/config.Config.GetScreenWidth()),
			rl.Color{v.R, v.G, v.B, v.A},
		)
	}

	r.console.Render()

	if config.Config.GetFPSEnabled() {
		rl.DrawFPS(10, 10)
	}

	rl.EndDrawing()
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

func (r RayLib) Run() {
	rl.InitWindow(int32(r.width), int32(r.height), "")
	rl.SetWindowTitle(config.Config.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	rl.SetConfigFlags(rl.FlagWindowResizable)

	for !rl.WindowShouldClose() {
		r.Update()
		r.Render()
	}
}
