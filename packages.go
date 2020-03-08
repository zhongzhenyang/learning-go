package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problem.\n", math.Sqrt(7))
}