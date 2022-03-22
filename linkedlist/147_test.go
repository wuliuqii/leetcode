package linkedlist

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2List([]int{4, 2, 1, 3})}, Ints2List([]int{1, 2, 3, 4})},
		{"regular", args{Ints2List([]int{-1, 5, 3, 4, 0})}, Ints2List([]int{-1, 0, 3, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
