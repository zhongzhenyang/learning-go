package main

import (
	"errors"
	"fmt"
	"os"
)

func _custom_error(name string) error {
	if len(name) == 0 {
		return errors.New("error:name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}

func _panic_error(index int) int {
	arr := [3]int{2,3,4}
	return arr[index]
}

func main() {
	_, err := os.Open("no_exist.txt")
	if err != nil {
		fmt.Println(err)
	}

	if err := _custom_error(""); err != nil {
		fmt.Println(err)
	}


	fmt.Println(_panic_error(5))

	fmt.Println("Can't occur")
}
