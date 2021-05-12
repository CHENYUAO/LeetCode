/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start

func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := make([]int, n)
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}
	//fmt.Printf("left: %v, right:%v", left, right)
	res := 0
	for i := 0; i < n; i++ {
		res += func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}(left[i], right[i])
	}
	return res
}

// @lc code=end

