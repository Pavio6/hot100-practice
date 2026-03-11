package tests

import (
	"testing"

	hot100practice "hot100_practice"
)

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{
			name: "valid board sample",
			board: [][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			want: true,
		},
		{
			name: "duplicate in row",
			board: [][]byte{
				[]byte("533.7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			want: false,
		},
		{
			name: "duplicate in column",
			board: [][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte("598....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			want: false,
		},
		{
			name: "duplicate in sub box",
			board: [][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8...5"),
			},
			want: false,
		},
		{
			name: "sparse but valid",
			board: [][]byte{
				[]byte("........."),
				[]byte("....1...."),
				[]byte("........."),
				[]byte("...2....."),
				[]byte("........."),
				[]byte(".....3..."),
				[]byte("........."),
				[]byte("....4...."),
				[]byte("........."),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.IsValidSudoku(tt.board)
			if got != tt.want {
				t.Fatalf("IsValidSudoku(%q) = %v, want %v", tt.board, got, tt.want)
			}
		})
	}
}
