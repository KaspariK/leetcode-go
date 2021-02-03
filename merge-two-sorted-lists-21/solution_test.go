package merge_two_sorted_lists_21

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := map[string]struct{
		listOne []int
		listTwo []int
		want	*ListNode
	}{
		"Test Case 1": {
			listOne: []int{1},
			listTwo: []int{2},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: nil,
				},
			},
		},
		"Test Case 2": {
			listOne: nil,
			listTwo: []int{2},
			want: &ListNode{
				Val: 2,
				Next: nil,
			},
		},
		"Test Case 3": {
			listOne: nil,
			listTwo: nil,
			want: nil,
		},
		"Test Case 4": {
			listOne: []int{1},
			listTwo: []int{0},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: nil,
				},
			},
		},
		"Test Case 5": {
			listOne: []int{1, 2, 3},
			listTwo: []int{1, 2, 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 3,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		"Test Case 6": {
			listOne: []int{2, 2, 3},
			listTwo: []int{1, 1, 5},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		a := SliceToLinkedList(tc.listOne)
		b := SliceToLinkedList(tc.listOne)

		got := mergeTwoLists(a, b)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}