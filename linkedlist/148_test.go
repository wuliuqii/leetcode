package linkedlist

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
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
		{"nil", args{Ints2List([]int{})}, Ints2List([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
