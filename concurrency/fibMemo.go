package main

import (
	"fmt"
)

func main() {
	numToCount := 100
	var a []uint64
	jobs := make(chan uint64, numToCount)
	results := make(chan uint64, numToCount)

	for i := 0; i < numToCount; i++ {
		jobs <- uint64(i)
		a = append(a, 0)
	}
	close(jobs)

	go worker(jobs, results, &a)
	go worker(jobs, results, &a)
	go worker(jobs, results, &a)
	go worker(jobs, results, &a)

	go worker(jobs, results, &a)
	go worker(jobs, results, &a)
	go worker(jobs, results, &a)
	go worker(jobs, results, &a)

	for j := 0; j < numToCount; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan uint64, results chan<- uint64, res *[]uint64) {
	for n := range jobs {
		results <- fib(uint64(n), res)
	}
}

func fib(n uint64, res *[]uint64) uint64 {
	if (*res)[n] != 0 {
		return (*res)[n]
	} else if n <= 1 {
		(*res)[n] = n
		return n
	}
	(*res)[n] = fib(n-1, res) + fib(n-2, res)
	return (*res)[n]
}
