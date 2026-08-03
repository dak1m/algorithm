package list

import (
	"reflect"
	"testing"
)

func Test_reverseListIteration(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: SimpleBuildListNode(1, 2, 3, 4, 5)},
			want: SimpleBuildListNode(5, 4, 3, 2, 1),
		},
		{
			args: args{head: SimpleBuildListNode(1, 2)},
			want: SimpleBuildListNode(2, 1),
		},
		{
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListIteration(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListRecursion(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: SimpleBuildListNode(1, 2, 3, 4, 5)},
			want: SimpleBuildListNode(5, 4, 3, 2, 1),
		},
		{
			args: args{head: SimpleBuildListNode(1, 2)},
			want: SimpleBuildListNode(2, 1),
		},
		{
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListRecursion(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
