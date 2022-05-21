package tree

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"regular", args{root: Ints2Tree([]int{2, 1, 3})}, true},
		{"regular", args{root: Ints2Tree([]int{5, 1, 4, NULL, NULL, 3, 6})}, false},
		{"same", args{root: Ints2Tree([]int{2, 2, 2})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
