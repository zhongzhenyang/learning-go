package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := f
	fmt.Println(x, y)
	fmt.Printf("z(%T):%v", z,z)


}
