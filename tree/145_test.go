package tree

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"regular", args{Ints2Tree([]int{1, NULL, 2, 3})}, []int{3, 2, 1}},
		{"nil", args{Ints2Tree([]int{})}, nil},
		{"single", args{Ints2Tree([]int{1})}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := postorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("postorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
