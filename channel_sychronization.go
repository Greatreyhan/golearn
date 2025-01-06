package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Starting")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	isDone := make(chan bool, 1)

	go worker(isDone)

	<-isDone
}
