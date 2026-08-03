package list

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbersII(t *testing.T) {
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
			args: args{
				l1: SimpleBuildListNode(7, 2, 4, 3),
				l2: SimpleBuildListNode(5, 6, 4),
			},
			want: SimpleBuildListNode(7, 8, 0, 7),
		},
		{
			args: args{
				l1: SimpleBuildListNode(0),
				l2: SimpleBuildListNode(0),
			},
			want: SimpleBuildListNode(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersII(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersII() = %v, want %v", got, tt.want)
			}
		})
	}
}
