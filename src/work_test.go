package main

import "testing"

func TestIsNumberDivisibleBy1000umber(t *testing.T) {

	t.Run("Test1 1", func(t *testing.T) {
		if !IsNumberDivisibleBy1000(1000, 1000) {
			t.Errorf("Expected true, got false")
		}
	})
}
