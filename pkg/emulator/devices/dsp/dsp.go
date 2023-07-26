package dsp

import (
	"strings"
	"time"

	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/dsp/note"
	"github.com/dfirebaugh/tortuga/pkg/emulator/devices/dsp/stream"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type DSP struct {
	player       *audio.Player
	audioContext *audio.Context
	Volume       float64
}

func (d *DSP) setAudioContext() {
	d.audioContext = audio.NewContext(44100)
}

func (d *DSP) SetVolume(v float64) {
	d.Volume = v
	d.player.SetVolume(v)
}

func (d *DSP) PlayNote(freq float32, duration time.Duration) {
	if d.audioContext == nil {
		d.setAudioContext()
	}
	if freq == 0 {
		time.Sleep(duration)
		return
	}
	p, _ := d.audioContext.NewPlayer(stream.New(d.audioContext.SampleRate(), freq, 15, 15))
	p.SetVolume(0.05)

	p.Play()
	time.Sleep(duration)
	p.Close()
}

func (d *DSP) Notes() map[string]float32 {
	return note.Notes
}

func (d *DSP) Frequency(letter string) float32 {
	return note.Notes[strings.ToUpper(letter)]
}

func (d *DSP) PlayNotes(notes []string, interval time.Duration) {
	frequencies := []float32{}
	for _, n := range notes {
		if freq, ok := note.Notes[strings.ToUpper(n)]; ok {
			frequencies = append(frequencies, freq)
		}
	}

	d.PlaySequence(frequencies, interval)
}

func (d *DSP) PlaySequence(sequence []float32, interval time.Duration) {
	for _, f := range sequence {
		d.PlayNote(f, interval)
	}
}

func (d *DSP) Update() {
}
