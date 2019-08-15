package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	// := 함수즉시할당

	// got := Hello("GODORI")
	// want := "Hello, GODORI"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// t.Run("saying hello to people", func(t *testing.T) {
	// 	got := Hello("Godori")
	// 	want := "Hello, Godori"

	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello, World"

	// 	assertCorrectMessage(t, got, want)
	// })

	t.Run("in Korean", func(t *testing.T) {
		got := Hello("고도리", "Korean")
		want := "안녕, 고도리"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sombra", "Spanish")
		want := "Hola, Sombra"
		assertCorrectMessage(t, got, want)
		// assert.Equal(t, a, b, "The two words should be the same.")
	})

	assert := assert.New(t)

	t.Run("in French", func(t *testing.T) {
		got := Hello("godori", "French")
		want := "Bonjour, godori!"
		assert.Equal(got, want, "test fail")
	})

}
