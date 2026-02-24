package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("Say Hello", message)
}

func main() {
	var wg sync.WaitGroup

	totalJobs := 5

	for i := 0; i < totalJobs; i++ {
		wg.Add(1)
		go sayHello(fmt.Sprintf("JOB %d", i), time.Second, &wg)
	}

	fmt.Println("Hello from Main() Goroutine")
	wg.Wait()
}
