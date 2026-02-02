package array

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "测试用例1",
			nums: []int{23, 2, 4, 6, 7},
			k:    6,
			want: true,
		},
		{
			name: "测试用例2",
			nums: []int{23, 2, 6, 4, 7},
			k:    6,
			want: true,
		},
		{
			name: "测试用例3",
			nums: []int{23, 2, 6, 4, 7},
			k:    13,
			want: false,
		},
		{
			name: "测试用例4",
			nums: []int{5, 0, 0, 0, 0},
			k:    3,
			want: true,
		},
		{
			name: "测试用例5",
			nums: []int{1, 2, 3},
			k:    5,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkSubarraySumByHashMap(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
