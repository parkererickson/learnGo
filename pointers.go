package main

import "fmt"

func main() {
	i := 7
	fmt.Println("Memory addres:", &i)

	inc1(i)
	fmt.Println("i not modified:", i)

	inc2(&i)
	fmt.Println("i incremented:", i)
}

func inc1(x int) {
	x++
}

func inc2(x *int) {
	*x++
}