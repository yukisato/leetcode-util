package datastructure

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_CreateListNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []int
		out *ListNode
	}{
		{nil, nil},
		{
			[]int{1},
			&ListNode{
				1,
				nil,
			},
		},
		{
			[]int{2, 2, 0},
			&ListNode{
				2,
				&ListNode{
					2,
					&ListNode{
						0,
						nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run("Test", func(t *testing.T) {
			t.Parallel()

			got := CreateListNode(tt.in)

			if !cmp.Equal(got, tt.out) {
				t.Errorf("got %#v want %#v", got, tt.out)
			}
		})
	}
}
