package main

import "fmt"

type Person interface {
	getName() string
}

type Man struct {
	name   string
	gender string
}

func (man *Man) getName() string {
	return man.name
}

type Woman struct {
	name   string
	gender string
}

func (woman *Woman) getName() string {
	return woman.name
}

func main() {
	var p1 Person = &Man{name: "Tom", gender: "male"}
	fmt.Println(p1.getName())

	//var _ Person = (*Woman)(nil)
	var p2 Person = &Woman{
		name:   "Alice",
		gender: "female",
	}
	w :=p2.(*Woman)
	fmt.Println(w.getName())
	//fmt.Println(p2.getName())
	fmt.Println(Woman{name: "Grace", gender: "female"})
}
