package leetcode0200

import "testing"

func TestNumIsLand(t *testing.T) {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{grid: [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, want: 6},
		{grid: [][]int{{1, 0, 1, 1}, {1, 0, 1, 0}, {0, 0, 0, 0}, {1, 0, 0, 1}}, want: 3},
	}
	for _, test := range tests {
		got := maxAreaOfIsland(test.grid)
		if got != test.want {
			t.Errorf("%v \n got: %d, want: %d\n", test.grid, got, test.want)
		}
	}
}
