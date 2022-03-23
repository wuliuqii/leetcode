package tree

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{Ints2Tree([]int{1, 2, 2, 3, 4, 4, 3})}, true},
		{"false", args{Ints2Tree([]int{1, 2, 2, NULL, 3, NULL, 3})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
