package ebiten

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type GameConsole interface {
	Update()
	Render(img *image.RGBA)
}

type Config interface {
	GetScreenHeight() int
	GetScreenWidth() int
	GetTitle() string
	GetScaleFactor() int
	GetFPSEnabled() bool
}

type Game struct {
	Width  int
	Height int
	ebiten.Game
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	Console         GameConsole
	img             *image.RGBA
	config          Config
}

func New(console GameConsole, c Config) *Game {
	return &Game{
		WindowTitle:     c.GetTitle(),
		WindowScale:     c.GetScaleFactor(),
		Width:           c.GetScreenWidth(),
		Height:          c.GetScreenHeight(),
		BackgroundColor: colornames.Skyblue,
		Console:         console,
		img:             image.NewRGBA(image.Rect(0, 0, c.GetScreenWidth(), c.GetScreenHeight())),
		config:          c,
	}
}
func (g *Game) Update() error {
	g.Console.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.Console.Render(g.img)
	screen.DrawImage(ebiten.NewImageFromImage(g.img), &ebiten.DrawImageOptions{})
	if g.config.GetFPSEnabled() {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func (g *Game) Reset(s interface{}) {
	g.Console = s.(GameConsole)
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
