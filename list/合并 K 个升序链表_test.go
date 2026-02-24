package list

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "mergeKLists",
			args: args{
				lists: SimpleBuildListNodes([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}),
			},
			want: SimpleBuildListNode(1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
