package shapes

import "math"

// Shape has to have an area method
type Shape interface {
	Area() float64
}

// Perimeter for rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Rectangle : float64, float64
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle : float64
type Circle struct {
	Radius float64
}

// Triangle : float64, float64
type Triangle struct {
	Base   float64
	Height float64
}

// Area for rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area for triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
