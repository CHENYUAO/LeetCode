/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0
	dump := &ListNode{
		Val : 0,
		Next : nil,
	}
	curr := dump
	for l1!=nil || l2!=nil {
		var n1 int = 0
		var n2 int = 0 
		if l1!=nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2!=nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1+n2+carry
		sum,carry = sum%10,sum/10
		curr.Next = &ListNode{
			Val : sum,
			Next : nil,
		}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val : carry,
			Next : nil,
		}
	}
	return dump.Next
}

// @lc code=end

