package linkedlist

import (
	"reflect"
	"testing"

	. "leetcode/tree"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"regular", args{Ints2List([]int{-10, -3, 0, 5, 9})}, Ints2Tree([]int{-0, -3, 9, -10, NULL, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
