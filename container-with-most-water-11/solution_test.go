package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := map[string]struct{
		heights []int
		want int
	}{
		"Basic": {heights: []int{1,8,6,2,5,4,8,3,7}, want: 49},
		"Short": {heights: []int{1, 1}, want: 1},
		"Largest Container": {heights: []int{2,3,4,5,18,17,6}, want: 17},
	}

	for name, tc := range tests {
		got := MaxArea(tc.heights)

		if tc.want != got {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}