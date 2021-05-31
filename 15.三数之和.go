/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
//Solution1：太丑陋了！每次i++和j--都要去重。。。
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	result := make([][]int, 0)
	sort.Ints(nums)
	for k := 0; k < n-2; k++ {
		target := -1 * nums[k]
		i, j := k+1, n-1
		for i < j {
			if nums[i]+nums[j] < target {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if nums[i]+nums[j] > target {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
		for k < n-1 && nums[k] == nums[k+1] {
			k++
		}
	}
	return result
}

//Solution2:超时┭┮﹏┭┮
// func threeSum(nums []int) [][]int {
// 	n := len(nums)
// 	if n < 3 {
// 		return [][]int{}
// 	}
// 	result := make([][]int, 0)
// 	sort.Ints(nums)
// 	for k := 0; k < n-2; k++ {
// 		target := -1 * nums[k]
// 		i, j := k+1, n-1
// 		for i < j {
// 			if nums[i]+nums[j] < target {
// 				i++
// 			} else if nums[i]+nums[j] > target {
// 				j--
// 			} else {
// 				temp := []int{nums[k], nums[i], nums[j]}
// 				if !haveExisted(result, temp) {
// 					result = append(result, temp)
// 				}
// 				i++
// 				j--
// 			}
// 		}
// 	}
// 	return result
// }
// func haveExisted(source [][]int, this []int) bool {
// 	for _, v := range source {
// 		if reflect.DeepEqual(v, this) {
// 			return true
// 		}
// 	}
// 	return false
// }

// @lc code=end

