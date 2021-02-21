package main

import "fmt"

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，
// 要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

func main() {
	test := Constructor()
	fmt.Println(test.Max_value())
	test.Push_back(1)
	test.Push_back(2)
	fmt.Println(test.Max_value())
	fmt.Println(test.Pop_front())
	fmt.Println(test.Max_value())

}

type MaxQueue struct {
	queue []int // 存储queue，
	max   []int // 最大队列

}

func Constructor() MaxQueue {
	return MaxQueue{make([]int, 0), make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.max) > 0 && this.max[len(this.max)-1] < value {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	if res == this.max[0] {
		this.max = this.max[1:]
	}
	return res
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
