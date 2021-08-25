/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

const INT_MAX = int(^uint(0) >> 1)

type MinStack struct {
	Stack    []int
	MinStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    []int{},
		MinStack: []int{INT_MAX},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if val < this.MinStack[len(this.MinStack)-1] {
		this.MinStack = append(this.MinStack, val)
	} else {
		this.MinStack = append(this.MinStack, this.MinStack[len(this.MinStack)-1])
	}
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

