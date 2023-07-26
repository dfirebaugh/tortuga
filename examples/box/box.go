package main

import (
	"image/color"

	"github.com/ByteArena/box2d"
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type Shape interface {
	Render()
}

type cart struct {
	tortuga.Console
	shapes []Shape
	world  *box2d.B2World
}

var (
	game          tortuga.Console
	pixelPerMeter = 2.0
)

type box struct {
	*box2d.B2Body
	color color.Color
}

func NewBox(world *box2d.B2World, x, y, w, h float64, density float64, friction float64, c color.Color) *box {
	boxBodyDef := box2d.MakeB2BodyDef()
	boxBodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	boxBodyDef.Position.Set(x, y)

	boxShape := box2d.MakeB2PolygonShape()
	boxShape.SetAsBox(w, h)
	boxFixtureDef := box2d.MakeB2FixtureDef()
	boxFixtureDef.Shape = &boxShape
	boxFixtureDef.Density = density
	boxFixtureDef.Friction = friction

	b := &box{
		world.CreateBody(&boxBodyDef),
		c,
	}

	b.CreateFixtureFromDef(&boxFixtureDef)

	return b
}

func (b *box) Update() {}
func (b box) Render() {
	fixture := b.GetFixtureList()
	if fixture == nil {
		return
	}
	boxShape := fixture.GetShape().(*box2d.B2PolygonShape)
	vertices := boxShape.M_vertices

	if len(vertices) < 2 {
		return
	}

	width := (vertices[1].X - vertices[0].X) * pixelPerMeter
	height := (vertices[2].Y - vertices[1].Y) * pixelPerMeter
	pos := b.GetPosition()

	rect := geom.MakeRect(
		(pos.X-width/2)*pixelPerMeter,
		float64(game.GetScreenHeight())-(pos.Y+(height/2))*pixelPerMeter,
		width,
		height,
	)
	geom.Draw(rect, game.GetDisplay(), b.color)
}

type boundary struct {
	*box2d.B2Body
	color color.Color
}

func NewBoundary(world *box2d.B2World, x, y, w, h float64, clr color.Color) *boundary {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position.Set(x, y)
	body := world.CreateBody(&bodyDef)

	b := box2d.MakeB2PolygonShape()
	b.SetAsBox(w, h)
	body.CreateFixture(&b, 0)

	return &boundary{
		body,
		clr,
	}
}

func (b *boundary) Update() {}
func (b boundary) Render() {
	fixture := b.GetFixtureList()
	if fixture == nil {
		return
	}
	boxShape := fixture.GetShape().(*box2d.B2PolygonShape)
	vertices := boxShape.M_vertices

	if len(vertices) < 2 {
		return
	}

	width := (vertices[1].X - vertices[0].X) * pixelPerMeter
	height := (vertices[2].Y - vertices[1].Y) * pixelPerMeter
	pos := b.GetPosition()

	rect := geom.MakeRect(
		pos.X*pixelPerMeter-width*pixelPerMeter/2,
		float64(game.GetScreenHeight())-(pos.Y+height/2)*pixelPerMeter,
		width*pixelPerMeter,
		height*pixelPerMeter,
	)
	geom.Fill(rect, game.GetDisplay(), b.color)
}

var count = 0

func (c *cart) Update() {
	c.world.Step(1.0/60.0, 8, 3)
	if count == 1000 {
		return
	}

	centerX := (float64(game.GetScreenWidth()) / pixelPerMeter) / 2
	topY := float64(game.GetScreenHeight()) / pixelPerMeter

	c.shapes = append(c.shapes, NewBox(
		c.world,
		centerX,
		topY,
		1,
		1,
		1,
		.3,
		game.Color(uint8(count%len(game.GetPalette()))),
	))
	count++
}

func (c cart) Render() {
	game.Clear()

	for _, s := range c.shapes {
		s.Render()
	}
}

func main() {
	game = tortuga.New()

	gravity := box2d.MakeB2Vec2(0, -9.8)
	world := box2d.MakeB2World(gravity)

	ground := NewBoundary(
		&world,
		0,
		1,
		float64(game.GetScreenWidth())/pixelPerMeter,
		1,
		game.Color(3),
	)
	left := NewBoundary(
		&world,
		((float64(game.GetScreenWidth()) / pixelPerMeter) / 3),
		float64(game.GetScreenHeight())/pixelPerMeter,
		2,
		float64(game.GetScreenHeight())/pixelPerMeter,
		game.Color(3),
	)
	right := NewBoundary(
		&world,
		((float64(game.GetScreenWidth())/pixelPerMeter)/3)*2,
		float64(game.GetScreenHeight())/pixelPerMeter,
		2,
		float64(game.GetScreenHeight())/pixelPerMeter,
		game.Color(3),
	)

	world.SetGravity(gravity)

	game.SetFPSEnabled(true)

	game.Run(&cart{
		game,
		[]Shape{
			ground,
			left,
			right,
		},
		&world,
	})
}
