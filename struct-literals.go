package main

import "fmt"

type Vertex3 struct {
	X, Y int
}

var (
	v1 = Vertex3{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex3{X: 1}  //Y:0 被隐式地赋予
	v3 = Vertex3{}      // X:0 Y:0
	p  = &Vertex3{1, 2} //创建一个 *Vertex3 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)

}
