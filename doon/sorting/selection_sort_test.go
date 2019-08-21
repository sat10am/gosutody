package sorting

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	t.Run("sort with numbers", func(t *testing.T) {
		numbers := []int{1, 23, 50, 24, 30, 99}
		got, e := SelectionSort(numbers)
		wanted := []int{1, 23, 24, 30, 50, 99}
		if e != nil {
			assert.Equal(t, got, wanted)
		}
	})
	t.Run("sort with empty numbers raise error", func(t *testing.T) {
		numbers := make([]int, 0)
		_, e := SelectionSort(numbers)
		wanted := errors.New(EmptySliceErrorMsg)
		assert.Equal(t, e, wanted)
	})
}
