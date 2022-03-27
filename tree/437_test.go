package tree

import "testing"

func Test_pathSum2(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			"regular",
			args{
				root:      Ints2Tree([]int{10, 5, -3, 3, 2, NULL, 11, 3, -2, NULL, 1}),
				targetSum: 8,
			},
			3,
		},
		{
			"regular",
			args{
				root:      Ints2Tree([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1}),
				targetSum: 22,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := pathSum2(tt.args.root, tt.args.targetSum); gotRes != tt.wantRes {
				t.Errorf("pathSum2() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
