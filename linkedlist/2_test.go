package linkedlist

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"regular",
			args{Ints2List([]int{2, 4, 3}), Ints2List([]int{5, 6, 4})},
			Ints2List([]int{7, 0, 8}),
		},
		{
			"regular",
			args{Ints2List([]int{9, 9, 9, 9, 9, 9, 9}), Ints2List([]int{9, 9, 9, 9})},
			Ints2List([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
