package leetcode0463

import "testing"

func TestIslandPerimeter(t *testing.T) {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{grid: [][]int{
			{1}}, want: 4},
		{grid: [][]int{
			{1, 0}}, want: 4},
		{grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, want: 16},
	}
	for _, test := range tests {
		got := islandPerimeter(test.grid)
		if got != test.want {
			t.Errorf("%v \n got: %d, want: %d\n", test.grid, got, test.want)
		}
	}
}
