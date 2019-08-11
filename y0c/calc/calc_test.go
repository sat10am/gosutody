package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	assert := assert.New(t)

	t.Run("operator '+' then return op1 + op2", func(t *testing.T) {
		got := Calc(3, 7, "+")
		want := 10

		assert.Equal(got, want, "got is not equal want")
	})

	t.Run("operator '-' then return op1 - op2", func(t *testing.T) {
		got := Calc(3, 7, "-")
		want := -4

		assert.Equal(got, want, "got is not equal want")
	})

	t.Run("operator '*' then return op1 * op2", func(t *testing.T) {
		got := Calc(3, 7, "*")
		want := 21

		assert.Equal(got, want, "got is not equal want")
	})

	t.Run("operator '/' then return op1 / op2", func(t *testing.T) {
		got := Calc(20, 5, "/")
		want := 4

		assert.Equal(got, want, "got is not equal want")
	})

	t.Run("unknown operator then return -1", func(t *testing.T) {
		got := Calc(20, 5, "c")
		want := -1

		assert.Equal(got, want, "got is not equal want")
	})
}
