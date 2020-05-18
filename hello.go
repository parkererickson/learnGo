package main

import (
	"fmt"
	"hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.Reverse("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
