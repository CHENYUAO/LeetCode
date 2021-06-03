/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	str := []byte(s)
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		if str[i] == str[i+1] {
			dp[i][i+1] = true
		}
	}
	res := make([]byte, 0)
	for l := 2; l < n; l++ {
		for i := 0; i+l < n; i++ {
			if str[i] == str[i+l] && dp[i+1][i+l-1] {
				dp[i][i+l] = true
			}
		}
	}
	//fmt.Printf("%v", dp)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] {
				if j-i+1 > len(res) {
					res = str[i : j+1]
				}
			}
		}
	}
	return string(res)
}

// @lc code=end

