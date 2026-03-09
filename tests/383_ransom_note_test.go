package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			name:       "leetcode sample 1",
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			name:       "leetcode sample 2",
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			name:       "leetcode sample 3",
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
		{
			name:       "empty ransom note",
			ransomNote: "",
			magazine:   "abc",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.CanConstruct(tt.ransomNote, tt.magazine)
			if got != tt.want {
				t.Fatalf("CanConstruct(%q, %q) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
			}
		})
	}
}
