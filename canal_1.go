package main

import "fmt"

func main() {
	canal := make(chan string)

	// T2
	go func() {
		canal <- "Golang Conference! Vindo da T2"
	}()

	// T1
	msg := <-canal // se o canal tiver cheio esvazia ele!!! MSG
	fmt.Println(msg)
}
