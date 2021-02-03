package merge_two_sorted_lists_21

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := map[string]struct{
		listOne *ListNode
		listTwo *ListNode
		want	*ListNode
	}{
		"Test Case 1": {
			listOne: &ListNode{
				Val: 1,
				Next: nil,
			},
			listTwo: &ListNode{
				Val: 2,
				Next: nil,
			},
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
			listTwo: &ListNode{
				Val: 2,
				Next: nil,
			},
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
			listOne: &ListNode{
				Val: 1,
				Next: nil,
			},
			listTwo: &ListNode{
				Val: 0,
				Next: nil,
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: nil,
				},
			},
		},
		"Test Case 5": {
			listOne: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: nil,
					},
				},
			},
			listTwo: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: nil,
					},
				},
			}, want: &ListNode{
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
			listOne: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: nil,
					},
				},
			},
			listTwo: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
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
		got := mergeTwoLists(tc.listOne, tc.listTwo)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}