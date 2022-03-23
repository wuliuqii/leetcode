package tree

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"regular", args{[]int{8, 5, 1, 7, 10, 12}}, Ints2Tree([]int{8, 5, 10, 1, 7, NULL, 12})},
		{"regular", args{[]int{1, 3}}, Ints2Tree([]int{1, NULL, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
