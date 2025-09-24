package leetcode0200

import "testing"

func TestNumIsLand(t *testing.T) {
	var tests = []struct {
		grid [][]byte
		want int
	}{
		{grid: [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}, want: 1},
		{grid: [][]byte{{1, 0, 1, 1}, {1, 0, 1, 0}, {0, 0, 0, 0}, {1, 0, 0, 1}}, want: 4},
	}
	for _, test := range tests {
		got := numIslands(test.grid)
		if got != test.want {
			t.Errorf("%v \n got: %d, want: %d\n", test.grid, got, test.want)
		}
	}
}
