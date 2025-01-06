package main

import "fmt"

func main() {
	queque := make(chan string, 2)
	queque <- "one"
	queque <- "two"
	close(queque)

	for elem := range queque {
		fmt.Println(elem)
	}
}
