package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

// important to note! assertion functions don't typically take pointers.
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
