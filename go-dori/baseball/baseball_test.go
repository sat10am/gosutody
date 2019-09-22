package main

import (
	"testing"
)

func TestBaseball(t *testing.T) {

	// assertCorrectMessage := func(t *testing.T, got, want int) {
	// 	t.Helper()
	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// }

	t.Run("Generate rand diff num in every round", func(t *testing.T) {
		got := generateRandNum()
		want := generateRandNum()

		t.Helper()
		if got == want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Check user answer and correct answer is equal", func(t *testing.T) {
		got := checkAnswer("0123", "0123")
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Display Scores when it mismatches number", func(t *testing.T) {
		userAnswer := "1235"
		correctAnswer := "1234"
		got := displayScore(userAnswer, correctAnswer)
		want := "o1 s3 b0 "

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
