package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestTwoSumII(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	want := []int{1, 2}

	got := hot100practice.TwoSumII(nums, target)
	// 使用 reflect.DeepEqual 来比较切片内容
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TwoSumII(%v, %d) = %v, want %v", nums, target, got, want)
	}
}
