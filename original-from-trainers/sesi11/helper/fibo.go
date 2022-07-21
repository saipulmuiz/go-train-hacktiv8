package helper

import "fmt"

/*
	buat sebuah function Fibonacci
	1. 	num = 10
		expected = "1,1,2,3,5,8"

	2.	num = 21
		expected = "1,1,2,3,5,8,13,21"
*/
func Fibonacci(num int) string {
	str := ""
	before := 0
	next := 0

	if num == 0 {
		return ""
	}

	for {
		if next == 0 {
			before = next
			next = 1
			str += fmt.Sprintf("%d,", next)
			continue
		}

		temp := next
		next += before
		if next > num {
			break
		}
		before = temp
		str += fmt.Sprintf("%d,", next)

	}

	str = string(str[:len(str)-1])

	return str
}
