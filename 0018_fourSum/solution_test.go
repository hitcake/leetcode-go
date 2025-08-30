package fourSum

import (
	"leetcode-go/util"
	"testing"
)

func TestFourSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
		{[]int{1, -2, -5, -4, -3, 3, 3, 5}, -11, [][]int{{-5, -4, -3, 1}}},
		{[]int{0, 0, 4, -2, -3, -2, -2, -3}, -1, [][]int{{-3, -2, 0, 4}}},
		{[]int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
	}
	for _, tt := range tests {
		got := fourSum(tt.nums, tt.target)
		if util.TwoDimArrayEqual(got, tt.want) == false {
			t.Errorf("expected: %v, got: %v", tt.want, got)
		}
	}
}
