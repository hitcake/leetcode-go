package threeSum

import (
	"leetcode-go/util"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var tests = []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0, 0, 0, 9}, [][]int{{0, 0, 0}}},
		{[]int{0, 1, 1}, [][]int{}},
	}
	for _, tt := range tests {
		got := threeSum(tt.nums)
		if util.TwoDimArrayEqual(got, tt.want) == false {
			t.Errorf("expected: %v, got: %v", tt.want, got)
		}
	}
}
