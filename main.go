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

	// Dynamically sized arrays
	length := 5
	capacity := 5
	x := make([]int, length, capacity)
	fmt.Println(x)

	// Appending to a slice
	// https://golang.org/pkg/builtin/#append
	x = append(x, 1, 2, 3)
	fmt.Println(x)
	// Test comment
}

