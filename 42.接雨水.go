/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
//dp
// func trap(height []int) int {
// 	n := len(height)
// 	if n == 0 {
// 		return 0
// 	}
// 	maxHeight := make([]int, n)
// 	maxHeight[0] = height[0]
// 	for i := 1; i < n; i++ {
// 		maxHeight[i] = max(maxHeight[i-1], height[i])
// 	}
// 	maxHeight[n-1] = min(height[n-1], maxHeight[n-1])
// 	for i := n - 2; i >= 0; i-- {
// 		maxHeight[i] = min(max(maxHeight[i+1], height[i]), maxHeight[i])
// 	}
// 	area := 0
// 	for i := 0; i < n; i++ {
// 		area += maxHeight[i] - height[i]
// 	}
// 	return area
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// stack
func trap(height []int) int {
	n := len(height)
	stack := make([]int, 0)
	area := 0
	begin := 0
	for begin < n-1 && height[begin] < height[begin+1] {
		begin++
	}
	for i := begin; i < n; i++ {
		// 出栈
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bot := height[stack[len(stack)-1]]
			for len(stack) > 0 && height[stack[len(stack)-1]] == bot {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				wid := i - left - 1
				if height[left]-bot < height[i]-bot {
					area += (height[left] - bot) * wid
				} else {
					area += (height[i] - bot) * wid
				}
			}
		}
		stack = append(stack, i)
	}
	return area
}

// @lc code=end

