package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Doon")

	got := buffer.String()
	want := "Hello, Doon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
