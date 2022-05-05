package tree

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"regular", args{Ints2Tree([]int{1, 2, 3, NULL, 5, NULL, 4})}, []int{1, 3, 4}},
		{"regular", args{Ints2Tree([]int{1, NULL, 3})}, []int{1, 3}},
		{"nil", args{Ints2Tree([]int{})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := rightSideView(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("rightSideView() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
