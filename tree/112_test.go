package tree

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"true",
			args{
				root:      Ints2Tree([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, NULL, 1}),
				targetSum: 22,
			},
			true,
		},
		{
			"false",
			args{
				root:      Ints2Tree([]int{1, 2, 3}),
				targetSum: 5,
			},
			false,
		},
		{
			"nil",
			args{
				root:      Ints2Tree([]int{}),
				targetSum: 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
