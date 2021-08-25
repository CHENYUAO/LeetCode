/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var curr int

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr = 0
	return recurList(head, n)
}

func recurList(currList *ListNode, n int) *ListNode {
	if currList == nil {
		return nil
	}
	currList.Next = recurList(currList.Next, n)
	curr++
	if curr == n {
		return currList.Next
	}
	return currList
}

// @lc code=end

