/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//右侧查找
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

// @lc code=end

