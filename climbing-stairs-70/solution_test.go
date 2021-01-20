package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := map[string]struct{
		input int
		want int
	}{
		"Test Case 1": {input: 2, want: 2},
		"Test Case 2": {input: 3, want: 3},
		"Test Case 3": {input: 4, want: 5},
		"Test Case 4": {input: 8, want: 34},
		"Test Case 5": {input: 1, want: 1},
	}

	for name, tc := range tests {
		got := ClimbStairs(tc.input)

		if got != tc.want {
			t.Errorf("%s; expected '%v', got '%v'", name, tc.want, got)
		}
	}
}