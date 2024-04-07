package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	go contador(10)
	go contador(10)
  contador(10)
}

func contador(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
