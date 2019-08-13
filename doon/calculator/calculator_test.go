package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	t.Run("Test add function", func(t *testing.T) {
		got := add(1, 2)
		wanted := 3
		assert.Equal(t, got, wanted)
	})
}
