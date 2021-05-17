/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	i, j := 0, n-1
	for i < j {
		curr := numbers[i] + numbers[j]
		if curr == target {
			return []int{i + 1, j + 1}
		} else if curr < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}

// @lc code=end

