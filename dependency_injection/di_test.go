package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nahua")

	got := buffer.String()
	want := "Hello, Nahua"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
