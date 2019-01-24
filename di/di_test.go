package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Evan")

	got := buffer.String()
	want := "Hello, Evan"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
