package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			args{Ints2List([]int{1, 2, 4}), Ints2List([]int{1, 3, 4})},
			Ints2List([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			"nil",
			args{Ints2List([]int{}), Ints2List([]int{})},
			Ints2List([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
