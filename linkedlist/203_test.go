package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2List([]int{1, 2, 6, 3, 4, 5, 6}), 6}, Ints2List([]int{1, 2, 3, 4, 5})},
		{"nil", args{Ints2List([]int{}), 1}, Ints2List([]int{})},
		{"duplicate", args{Ints2List([]int{7, 7, 7, 7}), 7}, Ints2List([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
