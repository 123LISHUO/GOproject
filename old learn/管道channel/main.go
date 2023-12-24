package main

import (
	"fmt"
	"sync"
)

func funone(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func funtwo(ch chan int) {

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {

	var ch = make(chan int, 10)
	wg.Add(1)
	funone(ch)
	wg.Add(1)
	funtwo(ch)
	wg.Wait()

}
