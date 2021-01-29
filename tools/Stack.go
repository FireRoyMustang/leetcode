package tools

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
