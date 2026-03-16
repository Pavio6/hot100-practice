package tests

import (
	"reflect"
	"testing"

	hot100practice "hot100_practice"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "leetcode sample 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "strictly increasing",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "strictly decreasing",
			temperatures: []int{60, 50, 40, 30},
			want:         []int{0, 0, 0, 0},
		},
		{
			name:         "all equal",
			temperatures: []int{70, 70, 70},
			want:         []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hot100practice.DailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("DailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
