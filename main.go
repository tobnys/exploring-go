package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello")

	fmt.Println("hello 2")
}
