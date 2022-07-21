package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main execution started")
	go firstProcess(8)
	secondProcess(8)

	time.Sleep(6 * time.Second)
	fmt.Println("active goroutine :", runtime.NumGoroutine())
	fmt.Println("main execution ended")
}

func display(msg string) {
	fmt.Println(msg)
}

func firstProcess(index int) {
	time.Sleep(5 * time.Second)
	fmt.Println("first process started ...")
	for i := 0; i < index; i++ {
		fmt.Println("first :", i+1)
	}
	fmt.Println("first process ended ...")
}

func secondProcess(index int) {
	fmt.Println("second process started ...")
	for i := 0; i < index; i++ {
		fmt.Println("second :", i+1)
	}
	fmt.Println("second process ended ...")
}
