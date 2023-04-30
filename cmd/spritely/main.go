package main

import (
	"flag"

	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/canvas"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/cursor"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/file"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/frames"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/inputcontroller"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/palette"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/toolbar"
	"github.com/dfirebaugh/tortuga/cmd/spritely/internal/widget"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/entity"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"github.com/dfirebaugh/tortuga/pkg/message/broker"
)

type cart struct {
	Game            tortuga.Console
	Palette         palette.Palette
	Cursor          cursor.Cursor
	Entities        []entity.Entity
	InputController inputcontroller.InputController
}

var spriteFile string

func (c cart) Render() {
	c.Game.Clear()
	geom.MakeRect(0, 0, float64(c.Game.GetScreenWidth()), float64(c.Game.GetScreenHeight())).
		Filled(c.Game.GetDisplay(), c.Game.Color(2))
	for _, e := range c.Entities {
		e.Render()
	}
}

func (c cart) Update() {
	for _, e := range c.Entities {
		e.Update()
	}
}

func initEntities(game tortuga.Console) []entity.Entity {
	b := broker.NewBroker()

	palette := &palette.Palette{
		Game:       game,
		MessageBus: b,
		Coordinate: component.Coordinate{
			X: 200,
			Y: 30,
		},
		Width:       4,
		ElementSize: 16,
	}
	cursor := &cursor.Cursor{
		Game:       game,
		MessageBus: b,
	}
	canvas := &canvas.Canvas{
		Game:       game,
		MessageBus: b,
		Width:      8,
		PixelSize:  16,
	}
	animationFrames := &frames.Frames{
		Game:       game,
		MessageBus: b,
		Coordinate: component.Coordinate{
			X: 0,
			Y: 160,
		},
	}
	var animations [][][]uint8
	if spriteFile != "" {
		animations = file.Load(spriteFile)
	}

	if len(animations) > 0 {
		animationFrames.Frames = animations[0]
	}

	toolbar := &toolbar.ToolBar{
		Game:       game,
		MessageBus: b,
		Coordinate: component.Coordinate{
			X: 0,
			Y: 130,
		},
		Width:     8,
		PixelSize: 3,
	}

	go b.Start()
	go cursor.Mailbox()
	go canvas.Mailbox()
	go palette.Mailbox()
	go animationFrames.Mailbox()

	entities := []entity.Entity{
		toolbar,
		palette,
		canvas,
		animationFrames,
		cursor,
	}
	// a widget is an entity that is interactive
	widgets := []widget.Widget{}

	// determine which of our entities are widgets
	// and add them to their own collection
	for _, e := range entities {
		if w, ok := e.(widget.Widget); ok {
			widgets = append(widgets, w)
		}
	}
	inputController := inputcontroller.InputController{
		Game:       game,
		Widgets:    widgets,
		MessageBus: b,
	}

	entities = append(entities, inputController)
	return entities
}

func main() {
	flag.StringVar(&spriteFile, "sprite", "", "pass in a filepath to open an existing sprite file")
	flag.Parse()

	game := tortuga.New()
	entities := initEntities(game)

	game.Run(cart{
		Game:     game,
		Entities: entities,
	})
}
