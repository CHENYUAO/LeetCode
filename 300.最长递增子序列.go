/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}
	res := 1
	// dp[i]：以nums[i]结尾的最长递增子序列
	//转移关系：dp[i] = 从0到i-1的所有dp，如果大于则+1取最大值
	//			否则为dp，同时和当前值取最大值
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
			res = max(res, dp[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

