package main

import "fmt"


func main() {
	i, j := 42, 2701
	p := &i // 指向 i
	fmt.Println(*p)  // 通过指针读取 i 的值
	*p = 21 // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j // 指向 j
	*p = *p /37 // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值

	num := 100
	_add(num)
	fmt.Println(num)
	_realAdd(&num)
	fmt.Println(num)
}

func _add(num int){
	num += 1
}

func _realAdd(num *int){
	*num +=1
}