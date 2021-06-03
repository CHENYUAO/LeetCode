/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
var dir []int = []int{-1, 0, 1, 0, -1}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		newx, newy := x+dir[i], y+dir[i+1]
		if newx >= 0 && newy >= 0 &&
			newx < len(grid) && newy < len(grid[0]) &&
			grid[newx][newy] == '1' {
			dfs(grid, newx, newy)
		}
	}
	return
}

// @lc code=end

