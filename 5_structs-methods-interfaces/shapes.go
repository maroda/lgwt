package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// the Interface is an "only what you need" design
// the Helper function can use it to pass arbitrary shape types
// so that the interface is where the type Rectangle,
// defined by 'rectangle', can be sent.
// the interface understands that the 'rectangle' variable being passed
// is a Rectangle type, which is satisfied by the Rectangle.Area method,
// so that is what is used to retrieve a return value.
// If the type passed to this interface was Circle, it would also find
// the satisfying method and use it. But if the type passed was Triangle,
// there is no matching method associated with a Triangle receiver for Area().
// The interface will fail to work.
type Shape interface {
	Area() float64
}

// type Rectangle has a method called Area that returns a float64
// so it satisfies the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// type Circle has a method that also returns float64 and satisfies Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// type Triangle method for satisfying Shape - it got here by adding a test first
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

/* Not used in the final version
// Simple function
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// This is the function Area(), it is not a method of anything,
// it does not have a receiver ("owner" of the method).
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
*/
