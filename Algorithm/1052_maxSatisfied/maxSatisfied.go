package main

import "fmt"

// 今天，书店老板有一家店打算试营业 customers.length 分钟。
// 每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，
// 那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
// 请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	fmt.Println(maxSatisfied([]int{10, 1, 7}, []int{0, 0, 0}, 3))
}
func maxSatisfied(customers []int, grumpy []int, X int) int {
	if X <= 0 {
		return sum(customers, grumpy)
	}
	res, tmp := 0, 0
	for j := 0; j < X; j++ {
		tmp += customers[j] * grumpy[j]
	}
	res = tmp
	for i := 0; i < len(grumpy)-X; i++ {
		tmp = tmp - customers[i]*grumpy[i] + customers[i+X]*grumpy[i+X]
		if tmp > res {
			res = tmp
		}
	}
	res += sum(customers, grumpy)
	return res
}
func sum(customers, grumpy []int) (res int) {
	for index, customer := range customers {
		if grumpy[index] == 0 {
			res += customer
		}

	}
	return
}
