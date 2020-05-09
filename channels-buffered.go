package main

import "fmt"

func main() {
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel: 
	// ch := make(chan int, 100)
	ch := make(chan int, 3)
	ch <- 1
	ch <- 4
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
