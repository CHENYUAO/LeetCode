/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
var dir []int = []int{-1, 0, 1, 0, -1}

func maxAreaOfIsland(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	m, n := len(grid), len(grid[0])
	var res int = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				curr := 0
				dfs(grid, i, j, &curr)
				res = max(res, curr)
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int, res *int) {
	(*res)++
	grid[x][y] = 0
	for i := 0; i < 4; i++ {
		newx, newy := x+dir[i], y+dir[i+1]
		if newx >= 0 && newy >= 0 && newx < len(grid) && newy < len(grid[0]) && grid[newx][newy] == 1 {
			dfs(grid, newx, newy, res)
		}
	}
}

// @lc code=end

