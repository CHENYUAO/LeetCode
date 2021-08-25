/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	// 左空、右空 返回true
	if left == nil && right == nil {
		return true
	}
	// 左右有一空 返回false
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	// 左右都不空，返回...
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

// @lc code=end

