package tests

import (
	hot100practice "hot100_practice"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	want := 6

	got := hot100practice.MaxSubArray(nums)
	if got != want {
		t.Fatalf("MaxSubArray(%v) = %d, want %d", nums, got, want)
	}
}
