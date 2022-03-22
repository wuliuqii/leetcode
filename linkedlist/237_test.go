package linkedlist

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"regular", args{Ints2List([]int{4, 5, 1, 9})}, Ints2List([]int{4, 1, 9})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node.Next)
			if !reflect.DeepEqual(tt.args.node, tt.want) {
				t.Errorf("deleteNode() = %v", tt.args)
			}
		})
	}
}
