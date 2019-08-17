package _struct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	t.Run("calculate perimeter of a rectangle", func(t *testing.T) {
		got := rectangle.Perimeter()
		wanted := 40.0
		assert.Equal(t, got, wanted)
	})
	t.Run("miscalculate perimeter of a rectangle", func(t *testing.T) {
		got := rectangle.Perimeter()
		wanted := 30.0
		assert.NotEqual(t, got, wanted)
	})
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	circle := Circle{10.0}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("calculate area of a rectangle", func(t *testing.T) {
		wanted := 72.0
		checkArea(t, rectangle, wanted)
	})

	t.Run("calculate area of a circle", func(t *testing.T) {
		wanted := 314.1592653589793
		checkArea(t, circle, wanted)
	})
}
