/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l < r {
		index := quickSelect(nums, k, l, r)
		if index == len(nums)-k {
			return nums[index]
		} else if index < len(nums)-k {
			l = index + 1
		} else {
			r = index - 1
		}
	}
	return nums[l]
}

func quickSelect(nums []int, k int, l, r int) int {
	temp := nums[l]
	i, j := l, r
	for i < j {
		for i < j && nums[j] >= temp {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= temp {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = temp
	return i
}

// @lc code=end

