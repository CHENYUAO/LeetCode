/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	//inorder
	curr := root
	for curr != nil || len(stack) > 0 {
		//push left
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		//visit
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		//push right
		if curr.Right != nil {
			curr = curr.Right
			//stack = append(stack, curr)
		} else {
			curr = nil
		}
	}
	return result
}

// @lc code=end

