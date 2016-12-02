package main

import (
	"fmt"
	"time"
	"math/rand"
)

func getRandom() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(10)
}

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", getRandom())
}
