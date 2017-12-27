package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getCurrentTime())
}

func getCurrentTime() int {
	t := time.Now()
	return t.Second()
}
