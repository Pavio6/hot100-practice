package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "true case", s: "abc", t: "ahbgdc", want: true},
		{name: "false case", s: "axc", t: "ahbgdc", want: false},
		{name: "empty s", s: "", t: "ahbgdc", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.IsSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("IsSubsequence(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
