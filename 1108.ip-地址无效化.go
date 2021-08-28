/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */

// @lc code=start
func defangIPaddr(address string) string {
	var res string
	for _, item := range address {
		if item == '.' {
			res += "[.]"
		} else {
			res += string(item)
		}

	}
	return res
}

// @lc code=end

