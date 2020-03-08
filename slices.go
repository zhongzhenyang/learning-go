package main

import (
	"fmt"
	"reflect"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:",reflect.TypeOf(primes)) //array
	var s  []int= primes[1:4] //一个半开区间，包括第一个元素，但排除最后一个元素。
	fmt.Println("s:", reflect.TypeOf(s)) //slice
	fmt.Println(s)
}
