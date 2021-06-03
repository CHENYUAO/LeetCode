/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
// func permute(nums []int) [][]int {
// 	n := len(nums)
// 	used := make([]bool, n)
// 	temp := make([]int, n)
// 	res := make([][]int, 0)

// 	backtrack(nums, temp, 0, &res, used)
// 	return res
// }

// func backtrack(nums, temp []int, curr int, res *[][]int, used []bool) {
// 	if curr == len(nums) {
// 		t := make([]int, len(nums))
// 		copy(t, temp)
// 		*res = append(*res, t)
// 		//fmt.Println(*res)
// 		return
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if !used[i] {
// 			used[i] = true
// 			temp[curr] = nums[i]
// 			backtrack(nums, temp, curr+1, res, used)
// 			temp[curr] = 0
// 			used[i] = false
// 		}
// 	}
// 	return
// }
//20210602
var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	n := len(nums)
	used := make([]bool, n)
	temp := make([]int, n)
	backtrack(nums, 0, temp, used)
	return result
}

func backtrack(nums []int, curr int, temp []int, used []bool) {
	if curr == len(nums) {
		t := make([]int, len(temp))
		copy(t, temp)
		result = append(result, t)
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			temp[curr] = nums[i]
			backtrack(nums, curr+1, temp, used)
			used[i] = false
			//temp[curr]
		}
	}
}

// @lc code=end

