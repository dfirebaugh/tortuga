package emulator

import (
	"image/color"
	"time"

	"github.com/dfirebaugh/tortuga/internal/emulator/devices/display"
)

type (
	configuration interface {
		GetTitle() string
		GetTileSize() int
		GetScaleFactor() int
		GetScreenHeight() int
		GetScreenWidth() int
		GetPalette() []color.Color
		Color(i uint8) color.Color
		RGBA(i uint8) color.RGBA
		GetVolume() int
		GetDebugEnabled() bool
		GetFPSEnabled() bool
		GetRenderPipelineDebug() bool
		GetColor(i uint8) color.Color
		SetRenderPipelineDebug(v bool)
		SetTitle(v string)
		SetScaleFactor(v int)
		SetPalette(v []color.Color)
		SetDebugEnabled(v bool)
		SetFPSEnabled(v bool)
		SetVolume(v int)
	}

	cart interface {
		Update()
		Render()
	}

	clock interface {
		GetTick() uint
	}

	fontProcessingUnit interface {
		PrintAt(s string, x int, y int, c uint8)
		ResetDisplay(display display.Displayer)
	}

	dsp interface {
		SetVolume(v float64)
		// PlaySequence plays a sequence of tones
		PlaySequence(sequence []float32, interval time.Duration)
		PlayNote(freq float32, duration time.Duration)
		PlayNotes(notes []string, interval time.Duration)
		// Frequency returns a frequency to match the note passed in
		//  e.g. Frequency("c3") returns 130.81
		Frequency(letter string) float32
		// Notes returns a map of note symbols and their matching frequencies
		Notes() map[string]float32
	}

	displayProvider interface {
		GetDisplay() display.Displayer
	}

	frameBuffer interface {
		display.Displayer
		displayProvider
		GetFrame() []byte
		Render()
	}
)
