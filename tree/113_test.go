package tree

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			"regular",
			args{
				root:      Ints2Tree([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1}),
				targetSum: 22,
			},
			[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			"nil",
			args{
				root:      Ints2Tree([]int{1, 2, 3}),
				targetSum: 5,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("pathSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
