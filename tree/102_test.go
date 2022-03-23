package tree

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{"regular", args{Ints2Tree([]int{3, 9, 20, NULL, NULL, 15, 7})}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"single", args{Ints2Tree([]int{3})}, [][]int{{3}}},
		{"nil", args{Ints2Tree([]int{})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := levelOrder(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("levelOrder() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
