package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, a := range arr {
		process1(i, a, 8)
		process2(i, a, 8)
		process3(i, a, 8)
	}

	last := time.Since(now)
	fmt.Println("time for execution :", last.Seconds())

}

func process1(index, num, find int) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 1 done... ")
		fmt.Println("find :", num)
	}
}

func process2(index, num, find int) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 2 done... ")
		fmt.Println("find :", num)
	}
}

func process3(index, num, find int) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 3 done... ")
		fmt.Println("find :", num)
	}
}
