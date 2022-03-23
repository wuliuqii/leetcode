package tree

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"regular",
			args{Ints2Tree([]int{1, 2, 5, 3, 4, NULL, 6})},
			Ints2Tree([]int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5, NULL, 6}),
		},
		{
			"single",
			args{Ints2Tree([]int{1})},
			Ints2Tree([]int{1}),
		},
		{
			"nil",
			args{Ints2Tree([]int{})},
			Ints2Tree([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want: %v", tt.args.root, tt.want)
			}
		})
	}
}
