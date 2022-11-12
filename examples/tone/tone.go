package main

import (
	"sort"
	"time"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/input"
)

type cart struct {
	input       input.PlayerInput
	game        tortuga.Console
	notes       []string
	currentNote int
}

func (c *cart) Update() {
	if c.input.IsUpJustPressed() {
		if c.currentNote == 0 {
			c.currentNote = len(c.notes)
		}
		c.currentNote = (c.currentNote - 1) % len(c.notes)
	}
	if c.input.IsDownJustPressed() {
		c.currentNote = (c.currentNote + 1) % len(c.notes)
	}
	if c.input.IsRightJustPressed() {
		c.game.PlayNote(c.game.Notes()[c.notes[c.currentNote]], time.Millisecond*250)
	}
	if c.input.IsLeftJustPressed() {
		go func() {
			c.game.PlayNote(c.game.Frequency("c4"), time.Second)
		}()
		go func() {
			c.game.PlayNote(c.game.Frequency("e4"), time.Second)
		}()
		go func() {
			c.game.PlayNote(c.game.Frequency("c4"), time.Second)
		}()
	}
}

func (c *cart) Render() {
	c.game.Clear()
	c.game.PrintAt(c.notes[c.currentNote], c.game.GetScreenWidth()/2, c.game.GetScreenHeight()/2, 6)
	c.game.PrintAt("press right arrow to play tone", 10, 180, 5)
	c.game.PrintAt("press up/down arrow to select a different note", 10, 195, 5)
}

func main() {
	game := tortuga.New()
	notes := []string{}
	for n := range game.Notes() {
		notes = append(notes, n)
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i] < notes[j]
	})

	game.Run(&cart{
		input: input.Keyboard{},
		game:  game,
		notes: notes,
	})
}
