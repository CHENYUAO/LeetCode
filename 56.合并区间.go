/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// fmt.Println(SortLeft(intervals))
	res := make([][]int, 0)
	SortLeft(intervals)
	// fmt.Println(intervals)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {

		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1] = []int{min(res[len(res)-1][0], intervals[i][0]),
				max(res[len(res)-1][1], intervals[i][1])}
		} else {
			res = append(res, intervals[i])
		}
		// fmt.Println(res)
	}
	return res
}

func SortLeft(intervals [][]int) {
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				temp := intervals[j]
				intervals[j] = intervals[i]
				intervals[i] = temp
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

