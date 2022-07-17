package main

import (
	"fmt"
	"time"
)

func main() {
	// -- Channels (Implementing channel) --
	// c := make(chan string)

	// go introduce("Airell", c)
	// go introduce("Nanda", c)
	// go introduce("Mailo", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// msg3 := <-c
	// fmt.Println(msg3)

	// close(c)

	// -- Channels (Channel with anonymous function) --
	// c := make(chan string)

	// students := []string{"Airell", "Mailo", "Indah"}

	// for _, v := range students {
	// 	go func(student string) {
	// 		fmt.Println("Student", student)
	// 		result := fmt.Sprintf("Hai, my name is %s", student)
	// 		c <- result
	// 	}(v)
	// }

	// for _, v := range students {
	// 	go introduce(v, c)
	// }

	// for i := 1; i <= 3; i++ {
	// 	print(c)
	// }

	// close(c)

	// -- UnBuffered Channel (Buffered vs unbuffered channel) --
	// c1 := make(chan int)

	// go func(c chan int) {
	// 	fmt.Println("func goroutine starts sending data into the channel")
	// 	c <- 10
	// 	fmt.Println("func goroutine after sending data into the channel")
	// }(c1)

	// fmt.Println("main goroutine sleeps for 2 seconds")
	// time.Sleep(time.Second * 2)

	// fmt.Println("main goroutine stars receiving data")
	// d := <-c1
	// fmt.Println("main goroutine received data:", d)

	// close(c1)
	// time.Sleep(time.Second)

	// -- Buffered Channel (Buffered vs unbuffered channel) --
	// c1 := make(chan int, 3)

	// go func(c chan int) {
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
	// 		c <- i
	// 		fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
	// 	}

	// 	close(c)
	// }(c1)

	// fmt.Println("main goroutine sleeps 2 seconds")
	// time.Sleep(time.Second * 2)

	// for v := range c1 { // v = <- c1
	// 	fmt.Println("main goroutine received value from channel", v)
	// }

	// -- Channel (Select) --
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}

func print(c <-chan string) {
	fmt.Println(<-c)
}
