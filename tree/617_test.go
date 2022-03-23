package tree

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"regular",
			args{
				root1: Ints2Tree([]int{1, 3, 2, 5}),
				root2: Ints2Tree([]int{2, 1, 3, NULL, 4, NULL, 7}),
			},
			Ints2Tree([]int{3, 4, 5, 5, 4, NULL, 7}),
		},
		{
			"regular",
			args{
				root1: Ints2Tree([]int{1}),
				root2: Ints2Tree([]int{2, 1}),
			},
			Ints2Tree([]int{3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
