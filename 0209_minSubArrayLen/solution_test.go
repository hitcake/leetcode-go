package minSubArrayLen

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	var tests = []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{11, []int{1, 1, 1, 1, 1}, 0},
	}
	for _, tt := range tests {
		got := minSubArrayLen(tt.target, tt.nums)
		if got != tt.want {
			t.Errorf("nums: %v, target: %d, want: %d, got: %d", tt.nums, tt.target, tt.want, got)
		}
	}
}
