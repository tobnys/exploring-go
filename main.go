package main

import (
	"fmt"
)

func main() {
	fmt.Println(swapStrings("Hello", "Hello2"))
}

func swapStrings(a, b string) (string, string) {
	return b, a
}
