package arraytest

import "testing"
import "reflect"

func TestSum(t *testing.T)  {

	assertCorrectMsg := func (t *testing.T, got int, want int, arr []int)  {
		t.Helper()
		if want !=  got {
			t.Errorf("got %d wnat %d given, %v", got, want, arr)
		}
	}

	t.Run("collection of 5 numbers", func (t *testing.T)  {
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		assertCorrectMsg(t, got, want, numbers)
	})

	t.Run("collection of any size", func (t *testing.T)  {
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6

		assertCorrectMsg(t, got, want, numbers)
	})

	t.Run("sum all test", func (t *testing.T)  {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3,9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum all tails", func (t *testing.T)  {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})

	checkSums := func(t *testing.T, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func (t *testing.T)  {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func (t *testing.T)  {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0,9}

		checkSums(t, got, want)
	})

}
