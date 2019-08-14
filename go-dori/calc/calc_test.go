package calc

import (
	"testing"
)

func TestCalc(t *testing.T){

	t.Run("Add Operator", func(t *testing.T)  {
		got := Add(1,2)
		want := 3
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	assertCorrectMessage := func(t *testing.T, got, want int){
		t.Helper()
		if got != want {
			t.Errorf("got %q wnat %q", got, want)
		}
	}

	t.Run("Subtract Operator", func (t *testing.T)  {
		got := Subtract(3,2)
		want := 1
		assertCorrectMessage(t, got, want)
	})

	t.Run("Multiply Operator", func(t *testing.T){
		got := Multiply(2,3)
		want := 6
		assertCorrectMessage(t, got, want)
	})

	t.Run("Devide Operator", func(t *testing.T){
		got := Devide(6,3)
		want := 2
		assertCorrectMessage(t, got, want)
	})

}