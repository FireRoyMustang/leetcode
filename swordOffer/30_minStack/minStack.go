package main

import "container/list"

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
// 调用 min、push 及 pop 的时间复杂度都是 O(1)。

type MinStack struct {
	stack    list.List
	minStack list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
	return minStack
}

func (this *MinStack) Push(x int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {

}

func (this *MinStack) Min() int {

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
