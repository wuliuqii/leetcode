package linkedlist

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2List([]int{1, 2, 3, 3, 4, 4, 5})}, Ints2List([]int{1, 2, 5})},
		{"regular", args{Ints2List([]int{1, 1, 1, 2, 3})}, Ints2List([]int{2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
