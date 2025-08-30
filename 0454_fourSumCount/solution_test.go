package fourSumCount

import "testing"

func TestFourSumCount(t *testing.T) {
	var tests = []struct {
		nums [][]int
		want int
	}{
		{[][]int{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}}, 2},
	}
	for _, test := range tests {
		got := fourSumCount(test.nums[0], test.nums[1], test.nums[2], test.nums[3])
		if got != test.want {
			t.Errorf("nums: %v expected: %v, got: %v", test.nums, test.want, got)
		}
	}
}
