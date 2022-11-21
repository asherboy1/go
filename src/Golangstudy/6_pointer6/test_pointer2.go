package main

import "fmt"

// 指向整型的指针
// p所传的就为地址 *取值 赋值 &取地址
// 存的变量为另一个变量的地址
func changeval1(p *int) {
	*p = 0
	fmt.Println("p = ", p)

}

func main() {
	//地址传递  改变了该地址中的值
	var a int = 10
	changeval1(&a)
	fmt.Println(a)
}
