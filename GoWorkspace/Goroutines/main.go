package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string)  {
	for i := 0; i < 5; i++  {
		fmt.Println(text)
		time.Sleep(400 * time.Millisecond)
	}
	channel <- "Message printed"
}

func main()  {
	var channel chan string
	go printMessage("Go is great!", channel)
	response := <- channel
	fmt.Println(response)
}