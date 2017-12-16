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
	var c int64 = 51231241234123123
	v := c
	fmt.Println(v)
}
