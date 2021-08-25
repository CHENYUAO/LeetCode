/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var res bool

func isPalindrome(head *ListNode) bool {
	res = true
	recursion(head, head)
	return res
}

func recursion(left, right *ListNode) {
	if right == nil {
		return
	}
	recursion(right.Next)
	if left.Val != right.Val {
		res = false
	}
	left = left.Next
}

// @lc code=end

