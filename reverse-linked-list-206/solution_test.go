package reverse_linked_list_206

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := map[string]struct {
		input *ListNode
		want  *ListNode
	}{
		"Test Case 1": {},
	}

	for name, tc := range tests {
		got := reverseList(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
