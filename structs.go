package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X int
	Y int
}

type Student struct {
	name string
	age  int
}

func (stu *Student) _hello(person string) string {
	return fmt.Sprintf("hello %s, Iam %s", person, stu.name)
}

func main() {
	fmt.Println(Vertex{1, 2})
	stu := &Student{
		name: "Tom",
	}
	println("stu:",reflect.TypeOf(stu))
	msg := stu._hello("Jack")
	fmt.Println(msg)

	stu2 := new(Student)
	fmt.Println(stu2._hello("Alice"))
}
