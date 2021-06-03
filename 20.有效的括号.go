/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)
	temp := []rune(s)
	for _, v := range temp {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch v {
			case ')':
				if pop != '(' {
					return false
				}
			case ']':
				if pop != '[' {
					return false
				}
			case '}':
				if pop != '{' {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// @lc code=end

