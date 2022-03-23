package tree

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"regular", args{Ints2Tree([]int{3, 9, 20, NULL, NULL, 15, 7})}, 2},
		{"regular", args{Ints2Tree([]int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6})}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
