package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Parker", age: 20}

	fmt.Println(p)
	fmt.Println(p.age)
}
