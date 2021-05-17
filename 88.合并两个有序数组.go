/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	curr := len(nums1) - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[curr] = nums1[i]
			i--
			curr--
		} else {
			nums1[curr] = nums2[j]
			j--
			curr--
		}
	}
	// for i >= 0 {
	// 	nums1[curr] = nums1[i]
	// 	i--
	// 	curr--
	// }
	for j >= 0 {
		nums1[curr] = nums2[j]
		j--
		curr--
	}
}

// @lc code=end

