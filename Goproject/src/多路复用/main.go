package main

import (
	"fmt"
)

func main() {
	var a = make(chan int, 10)
	var s = make(chan string, 10)
	for i := 0; i < 10; i++ {
		a <- i
	}

	for y := 0; y < 10; y++ {
		s <- "hello"
	}

	for {
		select {

		case v := <-a:
			fmt.Println(v, "is chan int")

		case v := <-s:
			fmt.Println(v, "is chan string")

		default:
			fmt.Println("over")
			return
		}
	}
}
