package leetcode0200

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			area := num(grid, i, j)
			if area > 0 {
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}

func num(grid [][]int, i, j int) int {
	if isNotInGrid(grid, i, j) {
		return 0
	}
	if isWater(grid, i, j) {
		return 0
	}
	if isVisited(grid, i, j) {
		return 0
	}
	grid[i][j] = grid[i][j] + 1
	return 1 + num(grid, i-1, j) + num(grid, i+1, j) + num(grid, i, j-1) + num(grid, i, j+1)
}

func isWater(grid [][]int, i, j int) bool {
	return grid[i][j] == 0
}
func isVisited(grid [][]int, i, j int) bool {
	return grid[i][j] > 1
}

func isNotInGrid(grid [][]int, i, j int) bool {
	return i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i])
}
