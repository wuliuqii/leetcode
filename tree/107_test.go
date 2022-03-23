package tree

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{"regular", args{Ints2Tree([]int{3, 9, 20, NULL, NULL, 15, 7})}, [][]int{{15, 7}, {9, 20}, {3}}},
		{"single", args{Ints2Tree([]int{3})}, [][]int{{3}}},
		{"nil", args{Ints2Tree([]int{})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := levelOrderBottom(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("levelOrderBottom() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
