package main

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func BenchmarkArea(b *testing.B) {
	for b.Loop() {
		Rectangle{5.0, 2.0}.Area()
		Circle{10.0}.Area()
		Triangle{10.0, 2.0}.Area()
	}
}

func ExampleRectangle_Area() {
	result := Rectangle{Width: 5.0, Height: 2.0}.Area()
	fmt.Println(result)
	// Output: 10
}

func ExampleCircle_Area() {
	result := Circle{Radius: 10.0}.Area()
	fmt.Println(result)
	// Output: 314.1592653589793
}

func ExampleTriangle_Area() {
	result := Triangle{Base: 10.0, Height: 2.0}.Area()
	fmt.Println(result)
	// Output: 10
}
