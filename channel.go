package main

import (
	"fmt"
	"time"
)

func SayHello(channel chan string) {
	fmt.Println("Initial SayHello")
	channel <- "Data for channels"
	fmt.Println("End Of SayHello")
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go SayHello(channel)

	fmt.Println("Channel Received! :", <- channel)

	fmt.Println("End OF Script")

	time.Sleep(500 * time.Millisecond)
}