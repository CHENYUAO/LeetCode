/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
	var sum int
	for {
		row := rand.Intn(7)
		col := rand.Intn(7)
		sum = row*7 + col
		if sum < 40 {
			break
		}
	}
	return sum%10 + 1
}

// @lc code=end

