package tree

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"true",
			args{
				p: Ints2Tree([]int{1, 2, 3}),
				q: Ints2Tree([]int{1, 2, 3}),
			},
			true,
		},
		{
			"false",
			args{
				p: Ints2Tree([]int{1, 2}),
				q: Ints2Tree([]int{1, NULL, 2}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
