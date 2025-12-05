package main

import "fmt"

var ch chan int
var done chan bool

func init() {
	ch = make(chan int, 1)
	done = make(chan bool)
}

func consumer() {
	//defer done<- true is not allowed
	defer func() { done <- true }()
	for {
		n := <-ch
		if n == -1 {
			break
		}
		fmt.Printf("Received:%d\n", n)
	}
}

func producer(n int) {
	defer func() { ch <- -1 }()
	for i := 0; i < n; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
	}
}

func main() {
	go consumer()
	go producer(5)
	<-done
}
