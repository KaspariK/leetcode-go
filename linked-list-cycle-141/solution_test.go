package linked_list_cycle_141

import "testing"
// TODO: clean this crap up
func TestHasCycle(t *testing.T) {
	// Test Case 1
	tc1Tail := &ListNode{
		Val: -4,
		Next: nil,
	}

	tc1Mid := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: tc1Tail,
		},
	}

	tc1Tail.Next = tc1Mid

	tc1Head := &ListNode{
		Val: 3,
		Next: tc1Mid,
	}

	// Test Case 2
	tc2Head := &ListNode{
		Val: 1,
		Next: nil,
	}

	tc2Tail := &ListNode{
		Val: 1,
		Next: tc2Head,
	}

	tc2Head.Next = tc2Tail

	// Test Case 3
	tc3Head := &ListNode{
		Val: 1,
		Next: nil,
	}

	tests := map[string]struct{
		input *ListNode
		want bool
	}{
		"Test Case 1": {input: tc1Head, want: true},
		"Test Case 2": {input: tc2Head, want: true},
		"Test Case 3": {input: tc3Head, want: false},
		"Test Case 4": {input: nil, want: false},
	}

	for name, tc := range tests {
		got := hasCycle(tc.input)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
