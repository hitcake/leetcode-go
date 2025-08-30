package util

import (
	"testing"
)

func TestSortTwoDimensionalArray(t *testing.T) {
	var tests = []struct {
		nums [][]int
		want [][]int
	}{
		{[][]int{{2, 1, 5}, {1, 2, 8}}, [][]int{{5, 1, 2}, {8, 1, 2}}},
	}
	for _, test := range tests {
		if TwoDimArrayEqual(test.nums, test.want) == false {
			t.Errorf("SortTwoDimensionalArray got: %v, want: %v", test.nums, test.want)
		}

	}
}
