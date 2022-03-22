package linkedlist

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2ListWithCycle([]int{3, 2, 0, -4}, 1)}, Ints2ListWithCycle([]int{2, 0, -4}, 0)},
		{"regular", args{Ints2ListWithCycle([]int{1, 2}, 0)}, Ints2ListWithCycle([]int{1, 2}, 0)},
		{"no ring", args{Ints2ListWithCycle([]int{1, 2}, -1)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
