package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "simple", s: "III", want: 3},
		{name: "with subtractive pair", s: "LVIII", want: 58},
		{name: "complex", s: "MCMXCIV", want: 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.RomanToInt(tt.s)
			if got != tt.want {
				t.Fatalf("RomanToInt(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
