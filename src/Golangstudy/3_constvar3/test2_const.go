package main

import "fmt"

func main() {

	// 常量 只读属性
	const a int = 10

	fmt.Println("a = ", a)

	// 枚举定义常量
	const (
		name = "tony"
		age  = 18
		sex  = "male"
	)

	const (
		// 可以在const中 添加关键字 iota 每行iota默认累加1，第一行默认为0  iota只能配合const

		zero = iota * 10
		first
		second
		thrid

		// 以第一个格式进行继承,变更就以变更为数据格式 iota不变
		forth = iota + 1
		fifth
	)

	fmt.Println("name = ", name)
	fmt.Println("second = ", second)
	fmt.Println("fifth = ", fifth)

}
