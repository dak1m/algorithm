package list

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				SimpleBuildListNode(1, 2, 2, 1),
			},
			want: true,
		},
		{
			args: args{
				SimpleBuildListNode(1, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeRecursion(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				SimpleBuildListNode(1, 2, 2, 1),
			},
			want: true,
		},
		{
			args: args{
				SimpleBuildListNode(1, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeRecursion(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
