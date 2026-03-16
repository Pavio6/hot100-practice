package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "single pair",
			s:    "()",
			want: true,
		},
		{
			name: "mixed valid",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "nested valid",
			s:    "{[]}",
			want: true,
		},
		{
			name: "mismatch",
			s:    "(]",
			want: false,
		},
		{
			name: "closing first",
			s:    "]",
			want: false,
		},
		{
			name: "leftover opening",
			s:    "(((",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.IsValid(tt.s)
			if got != tt.want {
				t.Fatalf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
