package tree

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"regular", args{root: Ints2Tree([]int{-10, 9, 20, NULL, NULL, 15, 7})}, 42},
		{"regular", args{root: Ints2Tree([]int{1, 2, 3})}, 6},
		{"regular", args{root: Ints2Tree([]int{2, -1})}, 2},
		{"regular", args{root: Ints2Tree([]int{9, 6, -3, NULL, NULL, -6, 2, NULL, NULL, 2, NULL, -6, -6, -6})}, 16},
		{"neg", args{root: Ints2Tree([]int{-3})}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxPathSum(tt.args.root); gotRes != tt.wantRes {
				t.Errorf("maxPathSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
