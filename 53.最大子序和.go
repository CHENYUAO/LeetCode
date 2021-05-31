/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	//dp
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	dp := make([]int, n, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		res = max(dp[i], res)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// @lc code=end

