package tree

import (
	"reflect"
	"testing"
)

func Test_buildTree2(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"regular",
			args{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			Ints2Tree([]int{3, 9, 20, NULL, NULL, 15, 7}),
		},
		{
			"single",
			args{[]int{3}, []int{3}},
			Ints2Tree([]int{3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree2(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree2() = %v, want %v", got, tt.want)
			}
		})
	}
}
