package main

import "fmt"

// 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
// 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	if dead("0000", deadends) {
		return -1
	}
	steps := 0
	queue := []string{"0000"}
	visted := make(map[string]bool)
	visted["0000"] = true
	size := len(queue)
	for size != 0 {
		for i := 0; i < size; i++ {
			str := queue[i]
			strs := adjustLock(str, deadends)
			for _, s := range strs {
				if s == target {
					return steps + 1
				}
				if visted[s] {
					continue
				} else {
					visted[s] = true
					queue = append(queue, s)
				}
			}
		}
		queue = queue[size:]
		size = len(queue)
		steps++
	}
	return -1
}
func adjustLock(lockNum string, deadends []string) []string {
	res := []string{}
	bytes := []byte(lockNum)
	for i := 0; i < 4; i++ {
		b := bytes[i]
		if b == '9' {
			bytes[i] = '0'
		} else {
			bytes[i] = b + 1
		}
		s := string(bytes)
		if !dead(s, deadends) {
			res = append(res, s)
		}
		if b == '0' {
			bytes[i] = '9'
		} else {
			bytes[i] = b - 1
		}
		s = string(bytes)
		if !dead(s, deadends) {
			res = append(res, s)
		}
		bytes[i] = b
	}
	return res
}

func dead(str string, deadends []string) bool {
	for _, dead := range deadends {
		if str == dead {
			return true
		}
	}
	return false
}
