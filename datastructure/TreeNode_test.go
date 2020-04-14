package datastructure

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_CreateTreeNodeFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		out *TreeNode
	}{
		{
			[]string{},
			nil,
		},
		{
			[]string{"0"},
			&TreeNode{
				0,
				nil,
				nil,
			},
		},
		{
			[]string{"1", "null", "2"},
			&TreeNode{
				1,
				nil,
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
		},
		{
			[]string{"3", "9", "20"},
			&TreeNode{
				3,
				&TreeNode{
					9,
					nil,
					nil,
				},
				&TreeNode{
					20,
					nil,
					nil,
				},
			},
		},
		{
			[]string{"3", "9", "20", "null", "null", "15", "7"},
			&TreeNode{
				3,
				&TreeNode{
					9,
					nil,
					nil,
				},
				&TreeNode{
					20,
					&TreeNode{
						15,
						nil,
						nil,
					},
					&TreeNode{
						7,
						nil,
						nil,
					},
				},
			},
		},
		{
			[]string{"0", "-3", "9", "-10", "null", "5"},
			&TreeNode{
				0,
				&TreeNode{
					-3,
					&TreeNode{
						-10,
						nil,
						nil,
					},
					nil,
				},
				&TreeNode{
					9,
					&TreeNode{
						5,
						nil,
						nil,
					},
					nil,
				},
			},
		},
		{
			[]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "null", "1"},
			&TreeNode{
				5,
				&TreeNode{
					4,
					&TreeNode{
						11,
						&TreeNode{
							7,
							nil,
							nil,
						},
						&TreeNode{
							2,
							nil,
							nil,
						},
					},
					nil,
				},
				&TreeNode{
					8,
					&TreeNode{
						13,
						nil,
						nil,
					},
					&TreeNode{
						4,
						nil,
						&TreeNode{
							1,
							nil,
							nil,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run("Test", func(t *testing.T) {
			t.Parallel()

			got := CreateTreeNodeFromString(tt.in)

			if !cmp.Equal(got, tt.out) {
				t.Errorf("got %#v want %#v", got, tt.out)
			}
		})
	}
}
