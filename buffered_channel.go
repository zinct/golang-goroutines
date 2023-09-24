package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	defer close(channel)

	channel <- "Indra Mahesa"
	fmt.Println(<- channel)
	channel <- "Mahesa"

	fmt.Println("End Of Script")

	time.Sleep(500 * time.Millisecond)

}