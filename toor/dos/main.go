package main

import "fmt"

// VisualElement that is drawn on the screen
type VisualElement interface {
	// Draw draws the visual element
	Draw(drawer *Drawer) error
}

// Square represents a square
type Square struct {
	// Location of the square
	Location Point
	// Side size
	Side float64
}

// Draw draws a square
func (square *Square) Draw(drawer *Drawer) error {
	return drawer.DrawRect(Rect{
		Location: square.Location,
		Size: Size{
			Height: square.Side,
			Width:  square.Side,
		},
	})
} // Circle represents a circle shape
type Circle struct {
	// Center of the circle
	Center Point
	// Radius of the circle
	Radius float64
}

// Draw draws a circle
func (circle *Circle) Draw(drawer *Drawer) error {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}
	return drawer.DrawEllipseInRect(rect)
}

type Layer struct {
	// Elements of visual elements
	Elements []VisualElement
}

// Draw draws a layer
func (layer *Layer) Draw(drawer *Drawer) error {
	for _, element := range layer.Elements {
		if err := element.Draw(drawer); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}
