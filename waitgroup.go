package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWg(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			workerWg(i)
		}()
	}

	wg.Wait()
}
