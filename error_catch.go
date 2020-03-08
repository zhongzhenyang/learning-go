package main

import "fmt"

func _catchError(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happend!", r)
			ret = -1
		}
	}()
	arr :=[3]int{2,3,4}
	return arr[index]
}

func main(){
	fmt.Println(_catchError(5))
	fmt.Println(_catchError(1))
	fmt.Println("here")
}
