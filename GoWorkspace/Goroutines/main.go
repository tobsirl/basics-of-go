package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string)  {
	for i := 0; i < 10; i++  {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 800)
	}
	channel <- "Message printed"
}

func main()  {
	var channel chan string
	go printMessage("Go is great!", channel)
	<- channel
	// go printMessage("Go is awesome!")
}