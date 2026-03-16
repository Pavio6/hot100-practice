package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "leetcode sample 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "leetcode sample 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "leetcode sample 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "single number",
			tokens: []string{"42"},
			want:   42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.EvalRPN(tt.tokens)
			if got != tt.want {
				t.Fatalf("EvalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
			}
		})
	}
}
