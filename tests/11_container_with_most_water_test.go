package tests

import (
	hot100practice "hot100_practice"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	want := 49

	got := hot100practice.MaxArea(height)
	if got != want {
		t.Fatalf("MaxArea(%v) = %d, want %d", height, got, want)
	}
}
