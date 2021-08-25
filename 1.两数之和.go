/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

// Solution 1
// func twoSum(nums []int, target int) []int {
// 	for index, num := range nums {
// 		for index2, num2 := range nums[index+1:] {
// 			if num+num2 == target {
// 				return []int{index, index + 1 + index2}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

// Solution 2
func twoSum(nums []int, target int) []int {
	// hash := map[int]int{}
	hash := make(map[int]int, 0)
	for index, num := range nums {
		p, ok := hash[target-num]
		if ok {
			return []int{p, index}
		}
		hash[num] = index
	}
	return []int{-1, -1}
}

// @lc code=end

