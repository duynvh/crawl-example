package main

import (
	"fmt"
	"time"
)

func main() {
	numberOfRequests := 100
	maxWorkerNumber := 5
	queueChan := make(chan int, numberOfRequests)
	doneChan := make(chan int)

	for i := 1; i <= maxWorkerNumber; i++ {
		go func(name string) {
			for v := range queueChan {
				crawl(name, v)
			}

			fmt.Printf("%s is done\n", name)
			doneChan <- i
		}(fmt.Sprintf("%d", i))
	}

	for i := 1; i <= numberOfRequests; i++ {
		queueChan <- i
	}

	close(queueChan)

	for i := 1; i <= maxWorkerNumber; i++ {
		<-doneChan
	}
}

func crawl(name string, v int) {
	time.Sleep(time.Second / 3)
	fmt.Printf("Worker %s is crawling: %d \n", name, v)
}
