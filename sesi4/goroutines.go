package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// -- Goroutine Example --
	// go goroutine()

	// -- Goroutines (Asynchronous process #1) --
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)

	time.Sleep(time.Second * 2)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println("main execution ended")
}

func goroutine() {
	fmt.Println("Hello")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}
