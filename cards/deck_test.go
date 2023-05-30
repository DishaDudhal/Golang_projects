package main

import "testing"

func TestNewDec(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got", len(d))
	}
}
