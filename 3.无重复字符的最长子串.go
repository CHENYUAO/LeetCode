/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	ss := []rune(s)

	l := 0
	res := 0
	for {
		r := l
		last := make(map[rune]int)
		for ; r < len(ss); r++ {
			_, ok := last[ss[r]]
			if ok {
				res = max(res, r-l)
				l = last[ss[r]] + 1
				break
			}
			last[ss[r]] = r
		}
		if r == len(ss) {
			res = max(res, r-l)
			break
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

