package main

import "testing"

func TestArea(t *testing.T) {

	// Anonymous slice of structs that will have the same format: shape, hasArea
	// then immediately filled with test values
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// Unnamed struct fields
		/*
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		*/
		// Optionally named struct fields are better in Testing
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	// step through the slice of structs
	// assign got to the current value's shape - which is an interface - to invoke the Area() method for that type
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}

/* This is the "non-table" version of the tests:

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	// this doesn't exist... the interface won't find this method
	// in fact that is exactly the error the test throws:
	// variable of Type Triangle cannot be passed to the Shape interface.
	// ./shapes_test.go:30:16: cannot use triangle (variable of type Triangle) as Shape value in argument to checkArea: Triangle does not implement Shape (missing method Area)
	// we don't "need" this, so we don't provide it through the interface
	// but the test is here to see the failure mode
	t.Run("triangles", func(t *testing.T) {
		triangle := Triangle{10, 5, 5}
		checkArea(t, triangle, 150)
	})
}
*/
