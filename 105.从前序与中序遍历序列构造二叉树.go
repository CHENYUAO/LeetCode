/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	return recursion(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func recursion(preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int) *TreeNode {
	// 递归出口
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	// 建root节点
	root := &TreeNode{
		Val: preorder[preStart],
	}
	// fmt.Println(root.Val)
	// 找该节点在树中的位置，inorder中该节点以左为左子树，以右为右子树
	var curr int
	for i, v := range inorder {
		if v == preorder[preStart] {
			curr = i
			break
		}
	}
	leftLen := curr - inStart
	root.Left = recursion(preorder, preStart+1, preStart+leftLen, inorder, inStart, curr-1)
	root.Right = recursion(preorder, preStart+leftLen+1, preEnd, inorder, curr+1, inEnd)
	return root
}

// @lc code=end

