/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int

const INT_MIN = ^int(^uint(0) >> 1)

func maxPathSum(root *TreeNode) int {
	res = INT_MIN
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	// 遍历到叶子节点，叶子节点值>0 返回其值，<0返回0
	if root == nil {
		return 0
	}
	fmt.Println("curr node:", root.Val)
	fmt.Println("curr res:", res)
	// 其他节点接收左右子树的返回值，
	left := dfs(root.Left)
	right := dfs(root.Right)
	//两个都>0，计算大的那个加上自己的值
	// curr := max(max(left, right), 0)
	//两个负，计算自己的值
	if left <= 0 && right <= 0 {
		res = max(res, root.Val)
		// res = max(root.Val, 0)
		if root.Val < 0 {
			return 0
		}
		return root.Val
	} else if left > 0 && right > 0 {
		if root.Val < 0 {
			res = max(res, max(left, right))
		} else {
			res = max(res, left+right+root.Val)
		}
		if root.Val+left+right < 0 {
			return 0
		}
		return root.Val + max(left, right)
	} else {
		// res = max(left+root.Val, right+root.Val)
		curr := max(left, right)
		if root.Val < 0 {
			res = max(res, curr)
		} else {
			res = max(res, curr+root.Val)
		}
		return curr + root.Val
	}
	//一正一负，计算正的那个加上自己的值
	// res = max(res, curr)
	// 计算值>0,返回，<0返回0
	// if curr+root.Val < 0 {
	// 	return 0
	// }
	// return curr + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

