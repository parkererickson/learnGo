package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[2] = 7
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	s := []int{5, 4, 3, 2, 1}
	fmt.Println(s)

	s = append(s, 0)
	fmt.Println(s)
}
