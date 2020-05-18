package main

import "fmt"

func main() {
	verticies := make(map[string]int)

	verticies["triangle"] = 3
	verticies["square"] = 4
	verticies["dodecagon"] = 12

	fmt.Println(verticies)
	fmt.Println(verticies["triangle"])

	delete(verticies, "square")

	fmt.Println(verticies)
}
