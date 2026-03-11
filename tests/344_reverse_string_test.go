package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want []byte
	}{
		{
			name: "odd length",
			in:   []byte("hello"),
			want: []byte("olleh"),
		},
		{
			name: "even length",
			in:   []byte("abcd"),
			want: []byte("dcba"),
		},
		{
			name: "single char",
			in:   []byte("a"),
			want: []byte("a"),
		},
		{
			name: "empty",
			in:   []byte{},
			want: []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hot100practice.ReverseString(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Fatalf("ReverseString() = %q, want %q", tt.in, tt.want)
			}
		})
	}
}
