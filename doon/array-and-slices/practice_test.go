package array_and_slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Run("Max function can find a largest value in integer slice", func(t *testing.T) {
		numbers := []int{55, 32, 40, 23, 24, 35, 65}
		got, err := Max(numbers)
		if err != nil {
			wanted := 65
			assert.Equal(t, got, wanted)
		}
	})

	t.Run("Max function returns error when called with empty slice", func(t *testing.T) {
		numbers := make([]int, 0)
		got, err := Max(numbers)
		wanted := -1
		assert.Error(t, err)
		assert.Equal(t, got, wanted)
	})
}

func TestMin(t *testing.T) {
	t.Run("Min function can find a largest value in integer slice", func(t *testing.T) {
		numbers := []int{55, 32, 40, 23, 24, 35}
		got, err := Min(numbers)
		if err != nil {
			wanted := 23
			assert.Equal(t, got, wanted)
		}
	})

	t.Run("Min function returns error when called with empty slice", func(t *testing.T) {
		numbers := make([]int, 0)
		got, err := Min(numbers)
		wanted := -1
		assert.Error(t, err)
		assert.Equal(t, got, wanted)
	})
}
