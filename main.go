package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // https://golang.org/pkg/time/#Time.UnixNano
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
