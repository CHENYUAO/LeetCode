/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	father := make(map[*TreeNode]*TreeNode)
	search(root, father)
	pathp := make(map[*TreeNode]bool)
	curr := p
	for curr != nil {
		pathp[curr] = true
		curr = father[curr]
	}
	curr = q
	for curr != nil {
		if _, ok := pathp[curr]; ok {
			return curr
		}
		curr = father[curr]
	}
	return nil
}

func search(root *TreeNode, father map[*TreeNode]*TreeNode) {
	if root.Left != nil {
		father[root.Left] = root
		search(root.Left, father)
	}
	if root.Right != nil {
		father[root.Right] = root
		search(root.Right, father)
	}
	return
}

// @lc code=end

