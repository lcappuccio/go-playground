package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("The time is", time.Now())
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
