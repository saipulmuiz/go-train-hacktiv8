package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	start := time.Now()
	fruits := []string{"Apel", "Jeruk", "Pisang", "Mangga", "Pepaya"}

	for i, fruit := range fruits {
		wg.Add(1)
		go printFruit(i, fruit, &wg)
	}

	wg.Wait()
	end := time.Since(start)
	fmt.Println("time execution :", end.Seconds())
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("index =>", index, "fruit =>", fruit)
	if wg != nil {
		wg.Done()
	}
}
