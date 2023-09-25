package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		defer close(channel)
		channel <- "Indra Mahesa"
		fmt.Println("End 1")
		channel <- "Indra Mahesa 2"
		fmt.Println("End 2")
	}()
	
	for data := range channel {
		fmt.Println("Menerima data:", data)
	}


	time.Sleep(500 * time.Millisecond)
}