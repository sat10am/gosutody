package array_and_slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("collections of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		wanted := 15
		assert.Equal(t, got, wanted)
	})

	t.Run("collections of any size numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Sum(numbers)
		wanted := 55

		assert.Equal(t, got, wanted)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	wanted := []int{3, 9}
	assert.Equal(t, got, wanted)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		wanted := []int{2, 9}
		assert.Equal(t, got, wanted)
	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		wanted := []int{0, 9}
		assert.Equal(t, got, wanted)
	})
}
