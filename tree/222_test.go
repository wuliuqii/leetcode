package tree

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"regular", args{Ints2Tree([]int{1, 2, 3, 4, 5, 6})}, 6},
		{"nil", args{Ints2Tree([]int{})}, 0},
		{"one", args{Ints2Tree([]int{1})}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
