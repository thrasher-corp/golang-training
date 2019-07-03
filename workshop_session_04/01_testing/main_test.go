package main

import "testing"

func TestPleaseReturnZero(t *testing.T) {
	if PleaseReturnZero() != 0 {
		t.Fatal("Oh dear!")
	}
}
