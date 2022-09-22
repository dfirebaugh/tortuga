package config

import (
	_ "embed"
	"image/color"
)

var (
	defaultPalette = []color.Color{
		color.Black,
		color.RGBA{127, 36, 84, 255},
		color.RGBA{28, 43, 83, 255},
		color.RGBA{0, 135, 81, 255},
		color.RGBA{171, 82, 54, 255},
		color.RGBA{96, 88, 79, 255},
		color.RGBA{195, 195, 198, 255},
		color.RGBA{255, 241, 233, 255},
		color.RGBA{237, 27, 81, 255},
		color.RGBA{250, 162, 27, 255},
		color.RGBA{247, 236, 47, 255},
		color.RGBA{93, 187, 77, 255},
		color.RGBA{81, 166, 220, 255},
		color.RGBA{131, 118, 156, 255},
		color.RGBA{241, 118, 166, 255},
		color.RGBA{252, 204, 171, 255},
	}
)

type config struct {
	Title string `yaml:"title"`
	// TileSize represents one dimension of a tile
	//   Tiles will always be square
	TileSize int `yaml:"tile-size"`
	Window   struct {
		Height int `yaml:"height"`
		Width  int `yaml:"width"`
	} `yaml:"window"`
	ScaleFactor  int  `yaml:"scale-factor"`
	DebugEnabled bool `yaml:"debug"`
	FPSEnabled   bool `yaml:"fps-enabled"`
	Volume       int  `yaml:"volume"`
	Palette      []color.Color
}

func Default() *config {
	return &config{
		Title: "tortuga",
		Window: struct {
			Height int `yaml:"height"`
			Width  int `yaml:"width"`
		}{Height: 240, Width: 320},
		ScaleFactor:  1,
		TileSize:     8,
		DebugEnabled: false,
		FPSEnabled:   false,
		Volume:       25,
		Palette:      defaultPalette,
	}
}

func (c config) GetTitle() string {
	return c.Title
}

func (c config) GetTileSize() int {
	return c.TileSize
}

func (c config) GetScaleFactor() int {
	return c.ScaleFactor
}
func (c config) GetScreenHeight() int {
	return c.Window.Height
}
func (c config) GetScreenWidth() int {
	return c.Window.Width
}
func (c config) GetPalette() []color.Color {
	return c.Palette
}
func (c config) GetVolume() int {
	return c.Volume
}
func (c config) GetDebugEnabled() bool {
	return c.DebugEnabled
}
func (c config) GetFPSEnabled() bool {
	return c.FPSEnabled
}

func (c *config) SetTitle(v string) {
	c.Title = v
}
func (c *config) SetTileSize(v int) {
	c.TileSize = v
}
func (c *config) SetScaleFactor(v int) {
	c.ScaleFactor = v
}
func (c *config) SetScreenHeight(v int) {
	c.Window.Height = v
}
func (c *config) SetScreenWidth(v int) {
	c.Window.Width = v
}
func (c *config) SetPalette(v []color.Color) {
	c.Palette = v
}
func (c *config) SetDebugEnabled(v bool) {
	c.DebugEnabled = v
}
func (c *config) SetFPSEnabled(v bool) {
	c.FPSEnabled = v
}
func (c *config) SetVolume(v int) {
	c.Volume = v
}
