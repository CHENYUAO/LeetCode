/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start

// 测试用例 151
//Testcase: "bbaa" "aba"
// Answer: "bba"
// Expected Answer:"baa"

//测试用例136
//Testcase: "aa"  "aa"
// Answer: "a"
// Expected Answer:"aa"

//本人代码，未AC，存在问题
// func containsEach(s []rune, t string) bool {
// 	var res bool = true
// 	for _, v := range []rune(t) {
// 		if !strings.ContainsRune(string(s), v) {
// 			res = false
// 		}
// 	}
// 	return res
// }

// func minWindow(s string, t string) string {
// 	if len(s) < len(t) {
// 		return ""
// 	}
// 	i, j := 0, len(t)
// 	var res []rune
// 	for j <= len(s) {
// 		temp := []rune(s[i:j])
// 		if containsEach(temp, t) {
// 			if len(res) == 0 || len(temp) < len(res) {
// 				//if len(temp) >= len(t) {
// 				res = temp
// 				//}
// 			}
// 			i++
// 			continue
// 		} else {
// 			j++
// 		}
// 	}
// 	return string(res)
// }

//答案代码
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// @lc code=end
