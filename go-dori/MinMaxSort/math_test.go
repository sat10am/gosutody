package math

import (
	"reflect"
	"testing"
)

func TestMath(t *testing.T) {

	assertCorrectMsg := func(t *testing.T, got, want interface{}) {
		switch got.(type) {
		case int:
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		case []int:
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}

	numbers := []int{3, 6, 2, 1, 5, 4}
	sortedNumbers := []int{1, 2, 3, 4, 5, 6}

	t.Run("Get Min Value from Array", func(t *testing.T) {
		got := Min(numbers)
		want := 1
		assertCorrectMsg(t, got, want)
	})

	t.Run("Get Max value from Array", func(t *testing.T) {
		got := Max(numbers)
		want := 6
		assertCorrectMsg(t, got, want)
	})

	t.Run("Sort the given array by bubble sorting", func(t *testing.T) {
		got := BubbleSort(numbers)
		want := sortedNumbers
		assertCorrectMsg(t, got, want)
	})

	t.Run("Sort the given array by Selection sorting", func(t *testing.T) {
		assertCorrectMsg(t, SelectionSort(numbers), sortedNumbers)
	})

}
