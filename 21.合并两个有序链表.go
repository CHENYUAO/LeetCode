/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	c1, c2 := l1, l2
	dump := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := dump
	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			curr.Next = c1
			c1 = c1.Next
		} else {
			curr.Next = c2
			c2 = c2.Next
		}
		curr = curr.Next
	}
	if c1 != nil {
		curr.Next = c1
	}
	if c2 != nil {
		curr.Next = c2
	}
	return dump.Next
}

// @lc code=end

