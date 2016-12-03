package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

// define a struc (record in C)
type Vertex struct {

	xCoordinate int
	yCoordinate int
}

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
	fmt.Println("My random  number is", getRandom())

	firstString, secondString := swapStrings("Swap", "strings")
	fmt.Println("Swap strings:", firstString, secondString)

	// Value conversions
	var myInt int = 42
	var myFloat float32 = float32(myInt) + math.Pi
	var myNewInt int = int(myFloat)
	fmt.Println(myInt, myFloat, myNewInt)

	// For loop (only loop in go!)
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i, sum)
		sum += i
	}
	fmt.Println("Total:", sum)

	// But the for becomes a while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Adding more:", sum)

	// An interesting variation
	for getRandom() < 5 {
		fmt.Println("Random is lesser than 5")
	}

	// Declare a struc
	var myVertex = Vertex{getRandom(),getRandom()}
	fmt.Println("Vertex is", myVertex)
	// And access its' fields
	fmt.Println("X coordinate is", myVertex.xCoordinate)
	fmt.Println("Y coordinate is", myVertex.yCoordinate)
}
