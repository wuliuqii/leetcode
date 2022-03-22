package linkedlist

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"regular", args{Ints2ListWithCycle([]int{3, 2, 0, -4}, 1)}, true},
		{"head", args{Ints2ListWithCycle([]int{3, 2}, 0)}, true},
		{"no ring", args{Ints2ListWithCycle([]int{3}, -1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
