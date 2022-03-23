package tree

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"regular",
			args{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			Ints2Tree([]int{3, 9, 20, NULL, NULL, 15, 7}),
		},
		{
			"single",
			args{[]int{3}, []int{3}},
			Ints2Tree([]int{3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
