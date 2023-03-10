package ebiten

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dfirebaugh/tortuga/config"
	"github.com/dfirebaugh/tortuga/internal/engine/game"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Game struct {
	Width  int
	Height int
	ebiten.Game
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	Console         game.Console
	img             *image.RGBA
	images          []*ebiten.Image
}

func New(console game.Console) *Game {
	g := &Game{
		WindowTitle:     config.Config.GetTitle(),
		WindowScale:     config.Config.GetScaleFactor(),
		Width:           config.Config.GetScreenWidth(),
		Height:          config.Config.GetScreenHeight(),
		BackgroundColor: colornames.Skyblue,
		Console:         console,
		img:             image.NewRGBA(image.Rect(0, 0, config.Config.GetScreenWidth(), config.Config.GetScreenHeight())),
		images:          []*ebiten.Image{},
	}

	g.updateRenderPipeline()
	return g
}
func (g *Game) Update() error {
	g.Console.Update()
	return nil
}

func (g *Game) shouldUpdateRenderPipeline() bool {
	return len(g.images) != len(g.Console.GetRenderPipeline())
}

func (g *Game) updateRenderPipeline() {
	g.images = []*ebiten.Image{}
	for _, img := range g.Console.GetRenderPipeline() {
		g.images = append(g.images, ebiten.NewImageFromImage(img))
	}
}

func (g *Game) renderRenderPipeline(screen *ebiten.Image) {
	for i, img := range g.Console.GetRenderPipeline() {
		op := &ebiten.DrawImageOptions{}
		eimg := g.images[i]
		op.GeoM.Translate(img.X, img.Y)
		screen.DrawImage(eimg, op)
	}
}

func (g *Game) renderDefaultFrameBuffer(screen *ebiten.Image) {
	screen.WritePixels(g.Console.GetFrame())
	g.Console.Render()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)

	if g.shouldUpdateRenderPipeline() {
		g.updateRenderPipeline()
	}

	g.renderDefaultFrameBuffer(screen)
	g.renderRenderPipeline(screen)

	if config.Config.GetFPSEnabled() {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
	}
	if config.Config.GetRenderPipelineDebug() {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("count: %d", len(g.Console.GetRenderPipeline())), 0, 10)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func (g *Game) Reset(s interface{}) {
	g.Console = s.(game.Console)
}

func (g *Game) Run() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc
		g.Exit()
	}()

	zoom := g.getZoom()
	ebiten.SetWindowSize(g.Width*zoom, g.Height*zoom)
	ebiten.SetWindowTitle(g.WindowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Exit() {
}

func (g *Game) getZoom() int {
	zoom := g.WindowScale
	if zoom == 0 {
		return 1
	}

	return zoom
}

func mergeLayers(frames ...[]uint8) []byte {
	result := frames[0]

	for _, f := range frames {
		copy(result, f)
	}
	return result
}
