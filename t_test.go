package main

import "testing"

func TestPrt(t *testing.T) {
	input := "Sun"
	expected := "Sun!"
	result := Prt(input)

	if result != expected {
		t.Errorf("Prt(%q) = %q; want %q", input, result, expected)
	}
}
