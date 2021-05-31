/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	lenA, lenB := 0, 0
	for a != nil {
		a = a.Next
		lenA++
	}
	for b != nil {
		b = b.Next
		lenB++
	}
	a, b = headA, headB
	if lenA > lenB {
		for i := 0; i < (lenA - lenB); i++ {
			a = a.Next
		}
	} else if lenA < lenB {
		for i := 0; i < (lenB - lenA); i++ {
			b = b.Next
		}
	} else {
		if headA == headB {
			return headA
		}
		for i := 0; i < lenA; i++ {
			a = a.Next
			b = b.Next
			if a == b {
				return a
			}
		}
		return nil
	}
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}

// @lc code=end

