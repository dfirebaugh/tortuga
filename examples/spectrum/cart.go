package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/ByteArena/box2d"
	"github.com/dfirebaugh/tortuga"
	"github.com/dfirebaugh/tortuga/pkg/component"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
)

type cart struct {
	tortuga.Console
	Player *Ship
	shapes []Shape
	world  *box2d.B2World
}

var toggle uint

var lastClick uint

func (c *cart) changeColor() {
	var clr color.Color
	toggle++
	if toggle >= 3 {
		toggle = 0
	}
	switch toggle {
	case 0:
		clr = game.Color(12)
	case 1:
		clr = game.Color(8)
	case 2:
		clr = game.Color(11)
	}
	c.Player.Color = clr
}

func (c *cart) Update() {
	if game.GamePad.IsLeftPressed() {
		c.Player.Rotation -= .1
	}
	if game.GamePad.IsRightPressed() {
		c.Player.Rotation += .1
	}
	if game.GamePad.IsUpPressed() {
		c.Player.Thrust += .01
	}
	if game.GamePad.IsDownPressed() {
		c.Player.Thrust -= .01
	}

	if game.GamePad.IsSecondaryJustPressed() {
		c.changeColor()
	}

	if game.Mouse.IsLeftClickJustPressed() {
		// if game.GetTick() < lastClick+100 {
		// 	println(lastClick)
		// 	c.changeColor()
		// }

		// lastClick = game.GetTick()
	}

	if game.Mouse.IsRightClickJustPressed() {
		c.changeColor()
	}

	if game.Mouse.IsRightClickPressed() {
		c.Player.Thrust += .008
	}

	if game.Mouse.IsLeftClickPressed() {
		const rotationThreshold = .1
		mx, my := game.Mouse.CursorPositionFloat()
		angle := geom.MakePoint(c.Player.X*pixelPerMeter, c.Player.Y*pixelPerMeter).ToVector().GetDirection(geom.MakePoint(mx, my).ToVector())
		rotationDiff := math.Abs(float64(c.Player.Rotation) - angle)
		if rotationDiff > rotationThreshold {
			c.Player.Rotation = component.Rotation(angle)
		}
		c.Player.Thrust += .008
	}

	c.Player.Update()
	i := 0
	for _, s := range c.shapes {
		if b, ok := s.(*Bullet); ok {
			if b.LifeTime > b.MaxLifeTime {
				// Remove the bullet from shapes
				c.shapes = append(c.shapes[:i], c.shapes[i+1:]...)
				continue
			}
		}
		if b, ok := s.(*Particle); ok {
			if b.LifeTime > b.MaxLifeTime {
				// Remove the bullet from shapes
				c.shapes = append(c.shapes[:i], c.shapes[i+1:]...)
				continue
			}
		}
		playerTri := geom.MakeTriangle([3]geom.Vector{
			geom.MakeVector(c.Player.X-c.Player.Width, c.Player.Y+c.Player.Height),
			geom.MakeVector(c.Player.X+c.Player.Width, c.Player.Y+c.Player.Height),
			geom.MakeVector(c.Player.X, c.Player.Y-c.Player.Height),
		})
		playerCir := geom.MakeCircle(playerTri.Centroid().X, playerTri.Centroid().Y, c.Player.Width)
		if o, ok := s.(*Orb); ok {
			oCir := geom.MakeCircle(o.X, o.Y, o.Width)
			if playerCir.HasOverlap(oCir) {
				// Remove the bullet from shapes
				c.shapes = append(c.shapes[:i], c.shapes[i+1:]...)
				if c.Player.Color != o.Color {
					// if o.Color != game.Color(8) || o.Color != game.Color(11) || o.Color != game.Color(12) {
					// 	continue
					// }
					c.explode(o, game.Color(4), 1.6)
					continue
				}

				var colorIndex ShipColor
				switch o.Color {
				case game.Color(8):
					c.explode(o, color.Black, -2)
					colorIndex = 0
				case game.Color(12):
					c.explode(o, color.Black, -2)
					colorIndex = 1
				case game.Color(11):
					c.explode(o, color.Black, -2)
					colorIndex = 2
				}
				if c.Player.lumenosity[colorIndex] < 1 {
					c.Player.lumenosity[colorIndex] += .01
				}
			}
		}
		i++
		s.Update()
	}
}

func (c *cart) explode(o *Orb, clr color.Color, thrust float64) {
	isMultiColor := clr == color.Black
	for i := 0; i < 20; i++ {
		pColor := clr

		if isMultiColor {
			pColor = game.Color(uint8(rand.Intn(16)))
		}
		c.shapes = append(c.shapes, &Particle{
			PhysicsObject: PhysicsObject{
				Position:  o.Position,
				Size:      component.Size{Height: .8},
				Friction:  .08,
				MaxThrust: 2,
				Thrust:    thrust,
				Rotation:  component.Rotation(rand.Intn(360)),
			},
			Color:       pColor,
			MaxLifeTime: 25,
		})
	}
}

func (c cart) Render() {
	game.Clear()
	// geom.Fill(
	// 	geom.MakeRect(
	// 		0,
	// 		0,
	// 		float64(game.GetScreenWidth()),
	// 		float64(game.GetScreenHeight())),
	// 	game.GetDisplay(),
	// 	colornames.Wheat)

	// geom.MakeLine(geom.MakePoint(c.Player.X*pixelPerMeter, c.Player.Y*pixelPerMeter), geom.MakePoint(game.Mouse.CursorPositionFloat())).Draw(game.GetDisplay(), game.Color(5))

	c.Player.Render()
	// game.RenderPalette()

	for _, s := range c.shapes {
		s.Render()
	}
}
