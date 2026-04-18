package main

import "testing"

func TestFtoC(t *testing.T) {
	celcius := ftoc(212)
	if celcius != 100 {
		t.Errorf("Expected 100, got %g", celcius)
	}

	celcius = ftoc(32)
	if celcius != 0 {
		t.Errorf("Expected 0, got %g", celcius)
	}
}
