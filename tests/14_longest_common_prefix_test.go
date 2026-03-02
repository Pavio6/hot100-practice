package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "leetcode sample",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "no common prefix",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "single string",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "empty input",
			strs: []string{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.LongestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Fatalf("LongestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}

