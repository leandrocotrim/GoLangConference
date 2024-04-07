package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", workerId, x)
	}
}


func main() {
	ch := make(chan int)
	workersCount := 100
	
	for i := 0; i < workersCount; i++ {
		go worker(i, ch)
	}
	
	for i := 0; i < 10000; i++ {
		ch <- i
	}
}
