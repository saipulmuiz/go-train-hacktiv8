package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, a := range arr {

		go process1(i, a, 8, c1)
		go process2(i, a, 8, c2)
		go process3(i, a, 8, c3)
	}

	select {
	case msg := <-c1:
		fmt.Println("c1", msg)
		break
	case msg := <-c2:
		fmt.Println("c2", msg)
		break
	case msg := <-c3:
		fmt.Println("c3", msg)
		break
	}

	last := time.Since(now)
	fmt.Println("time for execution :", last.Seconds())

}

func process1(index, num, find int, c chan string) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 1 done... ")
		fmt.Println("find :", num)
		c <- "process 1 done"
	}
}

func process2(index, num, find int, c chan string) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 2 done... ")
		fmt.Println("find :", num)
		c <- "process 2 done"
	}
}

func process3(index, num, find int, c chan string) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 3 done... ")
		fmt.Println("find :", num)
		c <- "process 3 done"
	}
}
