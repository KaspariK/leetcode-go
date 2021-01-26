package container_with_most_water_11

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := map[string]struct{
		heights []int
		want int
	}{
		"Test Case 1": {heights: []int{1,8,6,2,5,4,8,3,7}, want: 49},
		"Test Case 2": {heights: []int{1, 1}, want: 1},
		"Test Case 3": {heights: []int{2,3,4,5,18,17,6}, want: 17},
	}

	for name, tc := range tests {
		got := maxArea(tc.heights)

		if tc.want != got {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}