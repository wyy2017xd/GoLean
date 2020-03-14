package main

import (
	"fmt"
	"time"
)

func readChan(chanName <-chan int) {
	time.Sleep(10 * time.Second)
	for i := 0; i < 10; i++ {
		a := <-chanName
		fmt.Println("a = ", a)
	}
}
func writeChan(chanName chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
		chanName <- i
	}
}
func main() {
	var mychan = make(chan int, 10)
	go writeChan(mychan)
	go readChan(mychan)
	time.Sleep(30 * time.Second)
}
