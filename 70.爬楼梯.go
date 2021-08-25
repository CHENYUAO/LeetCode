/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp1 := 1
	dp2 := 2
	res := 0
	if n == 1 {
		return dp1
	}
	if n == 2 {
		return dp2
	}
	for i := 2; i < n; i++ {
		res = dp1 + dp2
		dp1 = dp2
		dp2 = res
	}
	return res
}

// @lc code=end

