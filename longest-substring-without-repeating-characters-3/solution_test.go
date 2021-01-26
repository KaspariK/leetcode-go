package longest_substring_without_repeating_characters_3

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct{
		items string
		want int
	}{
		"Test Case 1": {items: "kevin", want: 5},
		"Test Case 2": {items: "keevin", want: 4},
		"Test Case 3": {items: "pwfeewv", want: 4},
		"Test Case 4": {items: "ww", want: 1},
		"Test Case 5": {items: "abcabcbb", want: 3},
	}

	for name, tc := range tests {
		got := lengthOfLongestSubstring(tc.items)

		if tc.want != got {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
