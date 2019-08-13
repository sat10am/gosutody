package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	t.Run("Test add function", func(t *testing.T) {
		got := Add(1, 2)
		wanted := 3
		assert.Equal(t, got, wanted)
	})
	t.Run("Test subtract function", func(t *testing.T) {
		got := Subtract(100, 50)
		wanted := 50
		assert.Equal(t, got, wanted)
	})
}
