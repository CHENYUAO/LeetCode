/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	maxHeight := make([]int, n)
	maxHeight[0] = height[0]
	for i := 1; i < n; i++ {
		maxHeight[i] = max(maxHeight[i-1], height[i])
	}
	maxHeight[n-1] = min(height[n-1], maxHeight[n-1])
	for i := n - 2; i >= 0; i-- {
		maxHeight[i] = min(max(maxHeight[i+1], height[i]), maxHeight[i])
	}
	area := 0
	for i := 0; i < n; i++ {
		area += maxHeight[i] - height[i]
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

