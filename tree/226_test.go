package tree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"regular", args{Ints2Tree([]int{4, 2, 7, 1, 3, 6, 9})}, Ints2Tree([]int{4, 7, 2, 9, 6, 3, 1})},
		{"regular", args{Ints2Tree([]int{2, 1, 3})}, Ints2Tree([]int{2, 3, 1})},
		{"nil", args{Ints2Tree([]int{})}, Ints2Tree([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
