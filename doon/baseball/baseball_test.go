package baseball

import (
	"errors"
	"testing"
)

func validateAllCharactersAreDistinct(str string) (bool, error) {
	board := map[rune]bool{}
	for _, char := range str {
		_, exist := board[char]
		if exist {
			return false, errors.New("numbers are not unique")
		} else {
			board[char] = true
		}
	}
	return true, nil
}

func TestGenerateRandomNumber(t *testing.T) {
	got := GenerateRandomNumbers()

	t.Run("random number must be a 4-digit numbers", func(t *testing.T) {
		if len(got) != 4 {
			t.Errorf("got %s random numbers must be a 4-digit numbers", got)
		}
	})

	t.Run("each of random numbers are different", func(t *testing.T) {
		_, err := validateAllCharactersAreDistinct(got)
		if err != nil {
			t.Errorf("number %s is not distict", got)
		}
		nonUniqueString := "1111"
		_, err = validateAllCharactersAreDistinct(nonUniqueString)
		if err == nil {
			t.Errorf("nummber %s is invalid", nonUniqueString)
		}
	})
}

func TestAnswer(t *testing.T) {
	answer := NewAnswer()
	t.Run("new answer has not state", func(t *testing.T) {
		if answer.Ball != 0 {
			t.Errorf("new answer's ball count must be zero, got %d", answer.Ball)
		}
		if answer.Strike != 0 {
			t.Errorf("new answer's ball count must be zero, got %d", answer.Strike)
		}
		if answer.Out != 0 {
			t.Errorf("new answer's ball count must be zero, got %d", answer.Out)
		}
	})

	t.Run("each of new answer's numbers are unique", func(t *testing.T) {
		_, err := validateAllCharactersAreDistinct(answer.Numbers)
		if err != nil {
			t.Errorf("number %s is not distict", answer.Numbers)
		}
	})

	t.Run("set input value to answer", func(t *testing.T) {
		want := "1234"
		answer.SetInput(want)
		if answer.Input != want {
			t.Errorf("input value does not set, %s", want)
		}
	})
}
