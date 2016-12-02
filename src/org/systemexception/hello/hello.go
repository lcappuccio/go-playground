package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func getRandom() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(10)
}

func swapStrings(firstString, secondString string) (string, string) {
	return secondString, firstString
}

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", getRandom())

	firstString, secondString := swapStrings("Swap", "strings")
	fmt.Println("Swap strings: ", firstString, secondString)

	// Value conversions
	var myInt int = 42
	var myFloat float32 = float32(myInt) + math.Pi
	var myNewInt int = int(myFloat)
	fmt.Println(myInt, myFloat, myNewInt)
}
