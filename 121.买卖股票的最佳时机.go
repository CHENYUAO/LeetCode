/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
const INT_MAX = int(^uint(0) >> 1)

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	buyin := INT_MAX
	profiles := 0
	for i := 0; i < n; i++ {
		if prices[i] < buyin {
			buyin = prices[i]
		} else {
			profiles = max(profiles, prices[i]-buyin)
		}
	}
	return profiles
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

