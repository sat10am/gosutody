package StructInterface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	// got := Perimeter(10.0, 10.0)
	// rectangle := Rectangle{10.0, 10.0}
	// got := Perimeter(rectangle)
	// want := 40.0

	// if got != want {
	// 	t.Errorf("got %.2f want %.2f", got, want)
	// }
}

func TestArea(t *testing.T) {
	// === not use interface ===
	// rectangle := Rectangle{12.0, 6.0}
	// got := Area(rectangle)
	// want := 72.0

	// if got != want {
	// 	t.Errorf("got %.2f want %.2f", got, want)
	// }

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
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
}
