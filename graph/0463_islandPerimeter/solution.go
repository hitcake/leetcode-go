package leetcode0463

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				perimeter -= aroundIslandNum(grid, i, j)
			}

		}
	}
	return perimeter
}

func aroundIslandNum(grid [][]int, i, j int) int {
	islandNum := 0
	if isInGrid(grid, i-1, j) && isIsland(grid, i-1, j) {
		islandNum++
	}
	if isInGrid(grid, i+1, j) && isIsland(grid, i+1, j) {
		islandNum++
	}
	if isInGrid(grid, i, j-1) && isIsland(grid, i, j-1) {
		islandNum++
	}
	if isInGrid(grid, i, j+1) && isIsland(grid, i, j+1) {
		islandNum++
	}
	return islandNum
}

func isIsland(grid [][]int, i, j int) bool {
	return grid[i][j] == 1
}
func isInGrid(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i])
}
