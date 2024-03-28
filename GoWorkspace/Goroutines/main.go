package main

import (
	"fmt"
	"time"
)

func printMessage(text string)  {
	for i := 0; i < 10; i++  {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 800)
	}
}

func main()  {
	go printMessage("Go is great!")
	go printMessage("Go is awesome!")
}