package valid_palindrome_125

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test Case 1": {input: "Kevin", want: false},
		"Test Case 2": {input: "A man, a plan, a canal: Panama", want: true},
		"Test Case 3": {input: "race a car", want: false},
		"Test Case 4": {input: "racecar", want: true},
		"Test Case 5": {input: "a", want: true},
		"Test Case 6": {input: "", want: true},
		"Test Case 7": {input: "AbBa", want: true},
		"Test Case 8": {input: "ab", want: false},
		"Test Case 9": {input: "0P", want: false},
	}

	for name, tc := range tests {
		got := isPalindrome(tc.input)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
