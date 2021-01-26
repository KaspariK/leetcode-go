package reverse_bits_190

import (
	"testing"
)

func testReverseBits(t *testing.T) {
	tests := map[string]struct{
		input uint32
		want uint32
	}{
		"Test Case 1": {input: 43261596, want: 964176192},
		"Test Case 2": {input: 4294967293, want: 3221225471},
		"Test Case 3": {input: 2147483648, want: 1},
	}

	for name, tc := range tests {
		got := reverseBits(tc.input)

		if got != tc.want {
			t.Errorf("%s: ex[ected '%v', got '%v'", name, tc.want, got)
		}
	}
}