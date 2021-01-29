package main

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
// 调用 min、push 及 pop 的时间复杂度都是 O(1)。

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := MinStack{
		stack:    []int{1 << 31},
		minStack: []int{1 << 31},
	}
	return minStack
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	length := len(this.minStack)
	if x <= this.minStack[length-1] {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	if length > 1 {
		popNum := this.stack[length-1]
		this.stack = this.stack[:length-1]
		if popNum == this.Min() {
			this.minStack = this.minStack[:len(this.minStack)-1]
		}
	}

}

func (this *MinStack) Top() int {
	length := len(this.stack)
	if length > 1 {
		return this.stack[length-1]
	}
	return -1
}

func (this *MinStack) Min() int {
	length := len(this.minStack)
	if length > 1 {
		return this.minStack[length-1]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
