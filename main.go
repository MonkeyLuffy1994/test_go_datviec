package main

import (
	"fmt"
	"time"
)

const (
	numberOfURLs    = 1000
	numberOfWorkers = 5
)

func crawURL(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		time.Sleep(time.Second)
	}
	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {
	queue := make(chan int, 100)
	go func() {
		for i := 0; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}
		close(queue)
	}()
	return queue
}

func main() {
	var dataCombine []Hello
	var test = Hello{
		Name:  "Cong Danh",
		Email: "mail",
	}
	hellos := append(dataCombine, test)
	for _, item := range hellos {
		item.Name = "AAAAAAASDSDSDSDSDSDSDSDSDSDSDSD"
	}
	fmt.Printf(string(len(hellos)))
}
