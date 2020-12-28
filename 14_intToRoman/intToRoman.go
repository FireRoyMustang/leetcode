package main

import "fmt"

func main() {
	tests := []int{3, 5, 9, 1994}
	for _, test := range tests {
		fmt.Printf("%d:%s\n", test, intToRoman(test))
	}
}

func intToRoman(num int) string {
	a, b, digit := -1, 3, 1000
	ansBytes := make([]byte, 0)
	for b != -1 {
		a = num / digit
		num = num % digit
		if a != 0 {
			ansBytes = append(ansBytes, intToRomanChar(a, b)...)
		}
		b--
		digit /= 10
	}
	return string(ansBytes)
}

func intToRomanChar(num, digit int) []byte {
	romanChars := make([][]byte, 4)
	romanChars[0] = []byte{'I', 'V', 'X'}
	romanChars[1] = []byte{'X', 'L', 'C'}
	romanChars[2] = []byte{'C', 'D', 'M'}
	romanChars[3] = []byte{'M'}
	rets := make([]byte, 0)
	if num <= 3 {
		for i := 0; i < num; i++ {
			rets = append(rets, romanChars[digit][0])
		}
	} else if num == 4 {
		rets = append(rets, romanChars[digit][0])
		rets = append(rets, romanChars[digit][1])
	} else if num == 5 {
		rets = append(rets, romanChars[digit][1])
	} else if num <= 8 {
		rets = append(rets, romanChars[digit][1])
		for i := 5; i < num; i++ {
			rets = append(rets, romanChars[digit][0])
		}
	} else {
		rets = append(rets, romanChars[digit][0])
		rets = append(rets, romanChars[digit][2])
	}
	return rets
}
