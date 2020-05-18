package main

import (
	"fmt"
)

func main() {
	numToCount := 45
	jobs := make(chan int, numToCount)
	results := make(chan int, numToCount)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < numToCount; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < numToCount; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
