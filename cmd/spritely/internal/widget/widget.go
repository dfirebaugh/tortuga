package widget

import "github.com/dfirebaugh/tortuga/pkg/component"

// A widget is something that exists on
// the screen and has some level of interaction.
// It's capabable of translating screen coordinates to
// local coordinates within itself.
type Widget interface {
	// IsWithinBounds will determine if a coordinate exists within the widget.
	IsWithinBounds(coordinate component.Coordinate) bool
	// SelectElement will Set the element at the passed in coordinates to the active element.
	SelectElement(coordinate component.Coordinate)
	// AlternateSelectElement will Set the element at the passed in coordinates to the active element.
	AlternateSelectElement(coordinate component.Coordinate)
	// GetElement(coordinate component.Coordinate) Element
	Render()
	Update()
}

// type Element interface {
// 	// Render will display the Element at the given coordinates.
// 	Render(coordinate component.Coordinate)
// }
