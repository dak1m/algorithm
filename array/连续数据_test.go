package array

import "testing"

func Test_findMaxLengthAdvanced(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "测试用例1",
			nums: []int{0, 1, 1, 1, 1, 1, 0, 0, 0},
			want: 6, // [1,1,1,0,0,0]
		},
		{
			name: "测试用例2",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "测试用例3",
			nums: []int{0, 1, 0},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxLengthAdvanced(tt.nums)
			if got != tt.want {
				t.Errorf("findMaxLengthAdvanced() = %v, want %v", got, tt.want)
			}
			// 同时测试两个函数的结果是否一致
			got2 := findMaxLength(tt.nums)
			if got != got2 {
				t.Errorf("findMaxLengthAdvanced() = %v, findMaxLength() = %v, 结果不一致", got, got2)
			}
		})
	}
}
