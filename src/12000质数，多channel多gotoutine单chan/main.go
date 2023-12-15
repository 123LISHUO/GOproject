package main

import (
	"fmt"
	"sync"
	"time"
)

func putNum(InChan chan<- int) {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println("putNum", err)
		}

	}()

	for i := 2; i < 120000; i++ {
		InChan <- i
	}
	close(InChan)
	wg.Done()
}

func primeNum(InChan <-chan int, prinChan chan<- int, ExitChan chan<- bool) {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println("primeNum", err)
		}

	}()

	for nums := range InChan {
		var flag = true
		for i := 2; i < nums; i++ {
			if nums%i == 0 {
				flag = false
				break
			} else {
				flag = true
			}
		}
		if flag {
			// fmt.Println(nums,"is prime")
			prinChan <- nums
		}
	}
	ExitChan <- true
	wg.Done()

}
func printPrime(prinChan <-chan int) {
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("printPrime", err)
		}

	}()

	for num := range prinChan {
		fmt.Println(num, " is prime")
	}
	wg.Done()
}
func exitfun(ExitChan <-chan bool, prinChan chan<- int) {
	
	defer func (){

		if err := recover(); err != nil {
			fmt.Println("exitfun",err)
		}

	}()

	for i := 0; i < 60; i++ {
		<-ExitChan
	}
	close(prinChan)
	wg.Done()
}

var wg sync.WaitGroup

// var wx sync.Mutex
func main() {

	start:=time.Now().Unix()
	var InChan = make(chan int, 10000)
	var prinChan = make(chan int, 10000)
	var ExitChan = make(chan bool, 16)
	wg.Add(1)
	go putNum(InChan)
	for i := 1; i <= 60; i++ {
		wg.Add(1)
		go primeNum(InChan, prinChan, ExitChan)

	}
	wg.Add(1)
	go printPrime(prinChan)
	wg.Add(1)
	go exitfun(ExitChan, prinChan)

	wg.Wait()

	end:=time.Now().Unix()
	fmt.Println(end - start)
}
