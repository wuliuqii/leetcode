package linkedlist

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers2(t *testing.T) {
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
			args{Ints2List([]int{7, 2, 4, 3}), Ints2List([]int{5, 6, 4})},
			Ints2List([]int{7, 8, 0, 7}),
		},
		{
			"regular",
			args{Ints2List([]int{2, 4, 3}), Ints2List([]int{5, 6, 4})},
			Ints2List([]int{8, 0, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers2() = %v, want %v", got, tt.want)
			}
		})
	}
}
