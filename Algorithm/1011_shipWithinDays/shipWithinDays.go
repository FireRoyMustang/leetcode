package main

import "fmt"

// 传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
// 传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
// 返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}

func shipWithinDays(weights []int, D int) int {
	sumWeight, maxWeight := 0, 0
	for _, weight := range weights {
		sumWeight += weight
		maxWeight = max(maxWeight, weight)
	}
	if len(weights) == D {
		return maxWeight
	}
	left, right := maxWeight, sumWeight
	for left <= right {
		mid := (left + right) >> 1
		days := loadNeedDays(mid, D, weights)
		if days <= D {
			right = mid - 1
		} else if days > D {
			left = mid + 1
		}
	}
	return left
}
func loadNeedDays(load, D int, weights []int) int { //运载需要的天数
	days, loaded := 1, 0
	for i := 0; i < len(weights) && days <= D; i++ { // 注意 <=
		loaded += weights[i]
		if loaded > load {
			loaded = weights[i]
			days++
		}
	}
	// for _, weight := range weights {
	// 	loaded += weight
	// 	if loaded > load {
	// 		loaded = weight
	// 		days++
	// 		if days == D {
	// 			return days
	// 		}
	// 	}
	// }
	return days
}
func max(i, j int) int {
	if i < j {
		return j

	}
	return i
}
