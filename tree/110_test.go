package tree

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{Ints2Tree([]int{3, 9, 20, NULL, NULL, 15, 7})}, true},
		{"false", args{Ints2Tree([]int{1, 2, 2, 3, 3, NULL, NULL, 4, 4})}, false},
		{"nil", args{Ints2Tree([]int{})}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
