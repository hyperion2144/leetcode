package main

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				lists: []*ListNode{
					{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
					{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
					{Val: 2, Next: &ListNode{Val: 6}},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeKLists(tt.args.lists)
			for result != nil && tt.want != nil {
				if result.Val != tt.want.Val {
					t.Errorf("mergeKLists() = %v, want %v", result, tt.want)
					break
				}
				result = result.Next
				tt.want = tt.want.Next
			}
		})
	}
}
