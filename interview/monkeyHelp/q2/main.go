package main

import "fmt"

func main() {
	const MOD int = 1e9 + 7
	length := 0
	fmt.Scan(&length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&arr[i])
	}
	threshold := 0
	fmt.Scan(&threshold)
	res := 0
	lastNum, j := arr[0], 0
	if lastNum <= threshold {
		res++
		j++
	}
	for j < length {
		lastNum |= arr[j]
		if lastNum <= threshold {
			res++
			j++
		} else {
			break
		}
	}
	if j == length {
		tmp := res - 1
		for i := 1; i < length; i++ {
			res = (res + tmp) % MOD
			tmp--
		}
		fmt.Printf("%d\n", res)
		return
	}
	// j代表上一次到达的index+1
	for i := 1; i < length; i++ {
		if j > i { // j>=i代表 上一个数已经被|过了
			lastNum &= ^arr[i-1]
			res = (res + j - i - 1) % MOD
		} else {
			lastNum = 0
			j = i
		}
		for j < length {
			lastNum |= arr[j]
			if lastNum <= threshold {
				res = (res + 1) % MOD
				j++
			} else {
				break
			}
		}
	}
	fmt.Printf("%d\n", res)
}
