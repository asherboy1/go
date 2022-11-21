package main

import "fmt"

//声明变量四种方式

func main() {
	//申明变量 默认值为0
	var a int
	fmt.Println("a", a)
	fmt.Printf("%d aa", a)
	//
	var b int = 99
	fmt.Println("b", b)

	//最常用 :=
	c := "aabb"
	fmt.Println(c)

}
