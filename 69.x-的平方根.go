/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	//牛顿迭代法
	if x == 0 {
		return 0
	}
	c, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + c/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

// @lc code=end

