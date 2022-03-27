package tree

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"regular", args{Ints2Tree([]int{1, 2, 3})}, 25},
		{"regular", args{Ints2Tree([]int{4, 9, 0, 5, 1})}, 1026},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sumNumbers(tt.args.root); gotRes != tt.wantRes {
				t.Errorf("sumNumbers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
