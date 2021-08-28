/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	n := len(s)
	var res string
	if numRows == 1 {
		return s
	}
	dis := 2*(numRows) - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < n; j += dis {
			res += string(s[j])
			next := j + dis - i*2
			if i != 0 && i != numRows-1 && next < n {
				res += string(s[next])
			}
		}
	}
	return res
}

// @lc code=end

