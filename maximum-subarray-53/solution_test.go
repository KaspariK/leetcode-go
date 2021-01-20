package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := map[string]struct{
		input []int
		want int
	}{
		"Test Case 1": {input: []int{-2,1,-3,4,-1,2,1,-5,4}, want: 6},
		"Test Case 2": {input: []int{1}, want: 1},
		"Test Case 3": {input: []int{0}, want: 0},
		"Test Case 4": {input: []int{-1}, want: -1},
		"Test Case 5": {input: []int{-100000}, want: -100000},
		"Test Case 6": {input: []int{-2,-1,-3,-4,-1,-2,-1,-5,-4}, want: -1},
		"Test Case 7": {input: []int{0, 0, 0}, want: 0},
	}

	for name, tc := range tests {
		got := MaxSubArray(tc.input)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}