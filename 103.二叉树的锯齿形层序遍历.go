/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
    queue := make([]*TreeNode,0)
    res := make([][]int,0)
    if root == nil {
        return [][]int{}
    }
    cnt := 1
    queue = append(queue,root)
    for len(queue)>0 {
        length := len(queue)
        if cnt%2 == 0 {
            //偶数层从尾部出，从头部push
            this := make([]int,0)
            for i:=0;i<length;i++{
                // 出队列
                temp := queue[len(queue)-1]
                queue = queue[:len(queue)-1]   
                this = append(this,temp.Val) 
                if temp.Right != nil {
                    queue = append([]*TreeNode{temp.Right},queue...)
                }
                if temp.Left != nil {
                    queue = append([]*TreeNode{temp.Left},queue...)
                }
            }
            res = append(res,this)
        }else {
            //奇数层从头部出，从尾部push
            this := make([]int,0)
            for i:=0;i<length;i++{
                temp := queue[0]
                queue = queue[1:]
                this = append(this,temp.Val) 
                if temp.Left != nil{
                    queue = append(queue,temp.Left)
                }
                if temp.Right != nil {
                    queue = append(queue,temp.Right)
                }
            }
            res = append(res,this)
        }
        cnt++
    }
    return res	
}
// @lc code=end

