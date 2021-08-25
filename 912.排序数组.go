/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, begin, end int) {
	if begin > end {
		return
	}
	i, j := begin, end
	temp := nums[i]
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

	quickSort(nums, begin, i-1)
	quickSort(nums, i+1, end)
}

// @lc code=end

