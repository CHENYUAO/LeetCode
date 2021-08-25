/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
func pathInZigZagTree(label int) []int {
	row := findRow(label)
	res := []int{label}
	for i := row; i > 0; i-- {
		sum := int(math.Pow(2, i-1) + math.Pow(2, i) - 1)
		fmt.Println(sum)
		res = append(res, sum-(label/2))
	}
	return res
}

func findRow(label int) int {
	for row := 1; ; row++ {
		if math.Pow(2, row-1) <= label && label < math.Pow(2, row) {
			return int(row)
		}
	}
}

// @lc code=end

