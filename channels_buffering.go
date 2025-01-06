package main

import "fmt"

func main() {
	message := make(chan string, 2)

	go func() {
		message <- "ping1"
		message <- "ping2"
	}()

	fmt.Println(<-message)
	fmt.Println(<-message)
}
