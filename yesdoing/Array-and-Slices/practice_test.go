package main

import "testing"

func TestMin(t *testing.T) {

	t.Run("get Minimum value of Array", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}
		got := Min(numbers)
		want := 1

		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("get Maximum value of Array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Max(numbers)
		want := 5

		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
