package tree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"regular", args{Ints2Tree([]int{1, NULL, 2, 3})}, []int{1, 2, 3}},
		{"nil", args{Ints2Tree([]int{})}, nil},
		{"single", args{Ints2Tree([]int{1})}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := preorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("preorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
