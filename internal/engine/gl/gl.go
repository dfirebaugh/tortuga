package gl

import (
	"image"
	"image/color"
	"runtime"
	"time"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/internal/engine/game"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.org/x/image/colornames"
	"golang.org/x/image/draw"
)

type Game struct {
	Width           int
	Height          int
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	Console         game.Console
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var frames = 0.0
var last = 0.0

func New(emulator game.Console) game.Game {
	return &Game{
		Width:           config.Config.GetScreenWidth(),
		Height:          config.Config.GetScreenHeight(),
		WindowTitle:     config.Config.GetTitle(),
		WindowScale:     config.Config.GetScaleFactor(),
		BackgroundColor: colornames.Skyblue,
		Console:         emulator,
	}
}

func (g *Game) Run() {
	win := InitGLFW()
	defer Terminate()

	if win == nil {
		return
	}

	InitGL()
	InitFrameBuffer()

	for !win.ShouldClose() {
		g.render(win)
		go g.calculateFPS()
		go g.Console.Update()
		<-time.Tick(time.Second / 120)
	}
}

func (g *Game) Reset(interface{}) {

}

var img = image.NewRGBA(image.Rect(0, 0, config.Default().GetScreenWidth(), config.Default().Window.Height))
var scaledImg = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X*config.Default().ScaleFactor, img.Bounds().Max.Y*config.Default().ScaleFactor))

func (g *Game) render(win *glfw.Window) {
	c := config.Default()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// gl.UseProgram(prog)
	w, h := win.GetSize()
	g.Console.Render()
	gl.BindTexture(gl.TEXTURE_2D, Texture)

	// Set the expected size that you want:
	// Resize:
	draw.NearestNeighbor.Scale(scaledImg, scaledImg.Rect, img, img.Bounds(), draw.Over, nil)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(c.GetScreenWidth()*c.ScaleFactor), int32(c.GetScreenHeight()*c.ScaleFactor), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(imaging.FlipV(scaledImg).Pix))
	gl.BlitFramebuffer(0, 0, int32(w), int32(h), 0, 0, int32(c.GetScreenWidth()*c.ScaleFactor), int32(c.GetScreenHeight()*c.GetScaleFactor()), gl.COLOR_BUFFER_BIT, gl.LINEAR)

	glfw.PollEvents()
	win.SwapBuffers()
}

func (g *Game) calculateFPS() {
	// current := glfw.GetTime()
	// frames++
	// if current-last >= 1.0 {
	// 	// todo: add callback to setFPS
	// 	// g.config.SetFPS(frames)
	// 	frames = 0
	// 	last += 1.0
	// }
}
