/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	if n < 2 {
		if n == 0 {
			return -1
		} else {
			if nums[0] == target {
				return 0
			} else {
				return -1
			}
		}
	}
	for l < r {
		mid := (l + r) / 2
		//左侧升序
		if target == nums[mid] {
			return mid
		}
		if nums[mid] > nums[0] {
			if target >= nums[0] && target < nums[mid] {
				//左侧查找
				r = mid
			} else {
				//右侧查找
				l = mid + 1
			}
			//右侧升序
		} else if nums[mid] < nums[n-1] {
			if target > nums[mid] && target <= nums[n-1] {
				//右侧查找
				l = mid + 1
			} else {
				//左侧查找
				r = mid
			}
		} else {
			//nums[mid]==nums[0] || nums[mid]==nums[n-1]
			if mid == 0 {
				//you
				l = mid + 1
			} else { //mid==n-1
				r = mid
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

// @lc code=end

