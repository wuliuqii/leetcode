package tree

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"regular",
			args{
				root: Ints2Tree([]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}),
				p:    Ints2Tree([]int{5, 6, 2, NULL, NULL, 7, 4}),
				q:    Ints2Tree([]int{1, 0, 8}),
			},
			Ints2Tree([]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
