package longest_substring_without_repeating_characters_3

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct{
		items string
		want int
	}{
		"Test 1": {items: "kevin", want: 5},
		"Test 2": {items: "keevin", want: 4},
		"Test 3": {items: "pwfeewv", want: 4},
		"Test 4": {items: "ww", want: 1},
		"Test 5": {items: "abcabcbb", want: 3},
	}

	for name, tc := range tests {
		got := LengthOfLongestSubstring(tc.items)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
