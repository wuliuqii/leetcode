package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2List([]int{1, 2, 3, 4, 5}), 2}, Ints2List([]int{1, 2, 3, 5})},
		{"nil", args{Ints2List([]int{1}), 1}, Ints2List([]int{})},
		{"single", args{Ints2List([]int{1, 2}), 1}, Ints2List([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
