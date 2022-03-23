package tree

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"regular", args{Ints2Tree([]int{1, 2, 3, NULL, 5})}, []string{"1->2->5", "1->3"}},
		{"single", args{Ints2Tree([]int{1})}, []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := binaryTreePaths(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("binaryTreePaths() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
