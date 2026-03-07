package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{
			name:   "leetcode sample 1",
			jewels: "aA",
			stones: "aAAbbbb",
			want:   3,
		},
		{
			name:   "leetcode sample 2",
			jewels: "z",
			stones: "ZZ",
			want:   0,
		},
		{
			name:   "case sensitive",
			jewels: "a",
			stones: "AaaA",
			want:   2,
		},
		{
			name:   "empty stones",
			jewels: "abc",
			stones: "",
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.NumJewelsInStones(tt.jewels, tt.stones)
			if got != tt.want {
				t.Fatalf("NumJewelsInStones(%q, %q) = %d, want %d", tt.jewels, tt.stones, got, tt.want)
			}
		})
	}
}
