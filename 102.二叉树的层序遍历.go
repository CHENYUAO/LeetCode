/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root)
	//while(!queue.empty())
	for len(queue) > 0 {
		//for(int i=0;i<xx;++i)
		thisLen := len(queue)
		var temp []int
		for i := 0; i < thisLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			temp = append(temp, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

// @lc code=end

