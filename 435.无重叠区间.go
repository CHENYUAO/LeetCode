/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//fmt.Printf("%v", intervals)
	prev := intervals[0][1]
	res := 0
	for i := 1; i < n; i++ {
		if intervals[i][0] < prev {
			res++
		} else {
			prev = intervals[i][1]
		}
	}
	return res
}

// @lc code=end

