package tree

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"regular", args{Ints2Tree([]int{1, 3, 2, 5, 3, NULL, 9})}, 4},
		{"regular", args{Ints2Tree([]int{1, 3, NULL, 5, 3})}, 2},
		{"regular", args{Ints2Tree([]int{1, 3, 2, 5, NULL, NULL, 9, 6, NULL, NULL, 7})}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := widthOfBinaryTree(tt.args.root); gotRes != tt.wantRes {
				t.Errorf("widthOfBinaryTree() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
