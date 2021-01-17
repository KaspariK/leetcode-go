package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct{
		items []int
		target int
		want []int
	}{
		"Positive slice": {items: []int{46, 1562, 2244, 885, 2410, 3157, 1, 179, 4356}, target: 5241, want: []int{3, 8}},
	}

	for name, tc := range tests {
		got := TwoSum(tc.items, tc.target)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}