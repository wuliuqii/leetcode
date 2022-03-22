package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2List([]int{1, 2, 3, 4})}, Ints2List([]int{1, 4, 2, 3})},
		{"regular", args{Ints2List([]int{1, 2, 3, 4, 5})}, Ints2List([]int{1, 5, 2, 4, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if !reflect.DeepEqual(tt.args.head, tt.want) {
				t.Errorf("reorderList() = %v, want: %v", tt.args.head, tt.want)
			}
		})
	}
}
