package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	consume(ch)
}

func consume(ch chan int) {
	for x := range ch { // esvazia canal
		fmt.Println("Mensagem processada:", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}


