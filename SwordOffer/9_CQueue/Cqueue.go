package main

type CQueue struct {
	// stack1, stack2 *list.List
	stack1, stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if len(this.stack2) == 0 {
		len1 := len(this.stack1)
		for i := len1 - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = this.stack1[len1:]
	}
	if len(this.stack2) != 0 {
		e := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return e
	}
	return -1
}

// func Constructor() CQueue {
// 	return CQueue{
// 		stack1: list.New(),
// 		stack2: list.New(),
// 	}
// }

// func (this *CQueue) AppendTail(value int) {
// 	this.stack1.PushBack(value)
// }

// func (this *CQueue) DeleteHead() int {
// 	// 如果第二个栈为空
// 	if this.stack2.Len() == 0 {
// 		for this.stack1.Len() > 0 {
// 			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
// 		}
// 	}
// 	if this.stack2.Len() != 0 {
// 		e := this.stack2.Back()
// 		this.stack2.Remove(e)
// 		return e.Value.(int)
// 	}
// 	return -1
// }
