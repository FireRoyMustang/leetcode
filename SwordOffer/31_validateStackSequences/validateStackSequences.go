package main

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
// 假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
// 序列 {4,5,3,2,1}是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

func main() {

}
func validateStackSequences(pushed []int, popped []int) bool {
	stack := NewStack()
	ptr := 0
	for _, num := range pushed {
		stack.Push(num)
		for !stack.IsEmpty() && stack.Peek() == popped[ptr] {
			stack.Pop()
			ptr++
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]int, 0),
	}
}

func (this *Stack) Push(element int) {
	this.elements = append(this.elements, element)
}

func (this *Stack) IsEmpty() bool {
	return len(this.elements) == 0
}

func (this *Stack) Pop() int {
	length := len(this.elements)
	if this.IsEmpty() {
		panic("Pop error: Stack is empty.")
	}
	tmp := this.elements[length-1]
	this.elements = this.elements[:length-1]
	return tmp
}

func (this *Stack) Peek() int {
	if this.IsEmpty() {
		panic("Peek error: Stack is empty.")
	}
	return this.elements[len(this.elements)-1]
}
