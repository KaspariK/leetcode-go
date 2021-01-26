package valid_parentheses_20

import "testing"

func TestIsValid(t *testing.T) {
	tests := map[string]struct{
		input string
		want bool
	}{
		"Case 1": {input: "{}()[]", want: true},
		"Case 2": {input: "({})", want: true},
		"Case 3": {input: "({)}", want: false},
		"Case 4": {input: "(", want: false},
		"Case 5": {input: "[[", want: false},
		"Case 6": {input: "}", want: false},
		"Case 7": {input: "}}", want: false},
		"Case 8": {input: "(){}}{", want: false},
	}

	for name, tc := range tests {
		got := isValid(tc.input)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}