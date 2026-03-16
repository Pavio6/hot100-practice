package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "leetcode sample true",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "leetcode sample false",
			s:    "race a car",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "digits and letters",
			s:    "0P",
			want: false,
		},
		{
			name: "only symbols",
			s:    ".,",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.IsPalindrome(tt.s)
			if got != tt.want {
				t.Fatalf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
