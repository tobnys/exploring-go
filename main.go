package main

import (
	"fmt"
)

func main() {
	// Array
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// Slice
	var s []int = a[1:3]
	fmt.Println(s)

	d := a[1:3]
	fmt.Println(d)
}
