package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{head: Ints2List([]int{1, 2, 3, 4, 5})}, Ints2List([]int{5, 4, 3, 2, 1})},
		{"single", args{head: Ints2List([]int{1})}, Ints2List([]int{1})},
		{"nil", args{head: Ints2List([]int{})}, Ints2List([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
