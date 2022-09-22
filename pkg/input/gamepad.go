package input

type PlayerInput interface {
	IsUpPressed() bool
	IsDownPressed() bool
	IsLeftPressed() bool
	IsLeftJustPressed() bool
	IsRightPressed() bool
	IsRightJustPressed() bool
	IsPrimaryPressed() bool
	IsPrimaryJustPressed() bool
	IsSecondaryPressed() bool
	IsSecondaryJustPressed() bool
}

type Button uint8

const (
	Up Button = iota
	Down
	Left
	Right
	A
	B
)

type GamePad struct {
	Buttons map[Button]bool
	Device  PlayerInput
}

func (g *GamePad) Update() {
	g.Buttons[Up] = g.Device.IsUpPressed()
	g.Buttons[Down] = g.Device.IsDownPressed()
	g.Buttons[Left] = g.Device.IsLeftPressed()
	g.Buttons[Right] = g.Device.IsRightPressed()
	g.Buttons[A] = g.Device.IsPrimaryPressed()
	g.Buttons[B] = g.Device.IsSecondaryPressed()
}
