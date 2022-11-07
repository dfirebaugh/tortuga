package dsp

import (
	"time"

	"github.com/dfirebaugh/tortuga/internal/emulator/devices/dsp/stream"
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

func (d *DSP) PlayNote(freq int, duration time.Duration) {
	if d.audioContext == nil {
		d.setAudioContext()
	}
	if freq == 0 {
		time.Sleep(duration)
		return
	}
	p, _ := d.audioContext.NewPlayer(stream.New(d.audioContext.SampleRate(), freq))
	p.SetVolume(0.05)

	p.Play()
	time.Sleep(duration)
	p.Close()
}

func (d *DSP) PlaySequence(sequence []int, interval time.Duration) {
	for _, f := range sequence {
		d.PlayNote(f, interval)
	}
}

func (d *DSP) Update() {
}
