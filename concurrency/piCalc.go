package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	numThrows := 100000
	numWorkers := 8

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	numInRes := make(chan int, numWorkers)
	numOutRes := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			numToThrow := numThrows / numWorkers
			throw(numToThrow, numInRes, numOutRes)
			wg.Done()
		}()
	}
	wg.Wait()

	numIn := 0
	numOut := 0
	for i := 0; i < numWorkers; i++ {
		numIn = numIn + <-numInRes
		numOut = numOut + <-numOutRes
	}

	pi := 4 * (float64(numIn) / float64(numOut))
	fmt.Println("Pi is estimated to be:", pi)
}

func throw(numToThrow int, inside chan<- int, outside chan<- int) {
	in := 0
	out := 0

	for i := 0; i < numToThrow; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if math.Pow(x, 2)+math.Pow(y, 2) <= 1 {
			in++
		}
		out++
	}

	inside <- in
	outside <- out
}
