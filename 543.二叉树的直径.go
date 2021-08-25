/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
//Solution 1
// var res int

// func diameterOfBinaryTree(root *TreeNode) int {
// 	res = 0
// 	dfs(root)
// 	return res
// }

// // 返回值是节点到其子树的最大值,res是两节点直径的最大值
// func dfs(root *TreeNode) int {
// 	fmt.Println("current root:", root.Val)
// 	fmt.Println("current res:", res)
// 	//递归至叶子节点，叶子节点返回0
// 	if root.Left == nil && root.Right == nil {
// 		return 0
// 	}
// 	//左子树为空，返回右子树+1
// 	if root.Left == nil {
// 		right := 1 + dfs(root.Right)
// 		res = max(res, right)
// 		return right
// 	}
// 	//右子树为空，返回左子树+1
// 	if root.Right == nil {
// 		left := 1 + dfs(root.Left)
// 		res = max(res, left)
// 		return left
// 	}
// 	//左右子树都不为空，返回左右子树最大值+1
// 	left := 1 + dfs(root.Left)
// 	right := 1 + dfs(root.Right)
// 	res = max(res, left+right)
// 	return max(left, right)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

//Solution 2
var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) []int {
	//叶子节点返回[]int{0}
	// fmt.Println("current root:", root.Val)
	// fmt.Println("current res:", res)
	if root.Left == nil && root.Right == nil {
		return []int{0}
	}
	//递归 至 叶子节点
	if root.Left == nil {
		//遍历右子树
		rightLen := dfs(root.Right)
		for i := range rightLen {
			rightLen[i]++
			res = max(res, rightLen[i])
		}
		return rightLen
	}
	if root.Right == nil {
		//遍历左子树
		leftLen := dfs(root.Left)
		for i := range leftLen {
			leftLen[i]++
			res = max(res, leftLen[i])
		}
		return leftLen
	}
	//其他节点获取子节点的返回值，[]int
	//代表该节点到其子树的叶子节点的距离
	//此时直径最大值为两个数组最大值的和
	var (
		leftmax  int
		rightmax int
	)
	leftLen := dfs(root.Left)
	rightLen := dfs(root.Right)
	for i := range leftLen {
		leftLen[i]++
		leftmax = max(leftmax, leftLen[i])
	}
	for i := range rightLen {
		rightLen[i]++
		rightmax = max(rightmax, rightLen[i])
	}
	//返回两个数组拼接之后，所有值+1
	res = max(res, leftmax+rightmax)
	return append(leftLen, rightLen...)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

