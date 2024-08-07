package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer is a type in the bytes package that implements the Writer interface
	// because it has the method: Write(p []byte) (n int, err error)
	buffer := bytes.Buffer{}

	// Pass it as a reference to Greet, along with the string "Chris"
	Greet(&buffer, "Chris")
	// what we want to check is in the String method of the buffer type
	got := buffer.String()
	// and what we want from that String method is:
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
