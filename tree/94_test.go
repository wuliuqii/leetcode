package tree

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"regular", args{Ints2Tree([]int{1, NULL, 2, 3})}, []int{1, 3, 2}},
		{"nil", args{Ints2Tree([]int{})}, nil},
		{"single", args{Ints2Tree([]int{1})}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := inorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("inorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
