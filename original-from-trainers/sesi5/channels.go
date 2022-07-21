package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// students := []string{"Hacktiv8", "Koinworks", "Noobeeid"}

	// for _, student := range students {
	// 	go func(s string, ca chan string) {
	// 		fmt.Println(s)
	// 		str := fmt.Sprintf("Hello %s", s)
	// 		ca <- str
	// 	}(student, c)
	// }

	// for i := 0; i < 3; i++ {
	// 	print(c)
	// }

	go process1(c1)
	go process2(c2, c3)
	go process3(c3)

	for i := 0; i < 1; i++ {
		select {
		case msg := <-c1:
			fmt.Println("c1", msg)
			break
		case msg := <-c2:
			fmt.Println("c2", msg)
			break
		}
	}

}

func process1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "process 1 done"
}

func process2(c chan string, c3 chan string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("process from c3")
	msg := <-c3
	fmt.Println("process :", msg)
	c <- "process 2 done"
}

func process3(c chan string) {
	time.Sleep(time.Millisecond * 1200)
	c <- "process 3 done"
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(name string, c chan<- string) {
	str := fmt.Sprintf("Hello %s", name)

	c <- str
}
