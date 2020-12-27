package main

func main() {

}
func minArray(numbers []int) int {
	length := len(numbers)
	left, right := 0, length-1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return numbers[mid]

}
