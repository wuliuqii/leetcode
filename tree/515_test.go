package tree

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"regular", args{root: Ints2Tree([]int{1, 3, 2, 4, 3, NULL, 9})}, []int{1, 3, 9}},
		{"regular", args{root: Ints2Tree([]int{1, 2, 3})}, []int{1, 3}},
		{"nil", args{root: Ints2Tree([]int{})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := largestValues(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("largestValues() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
