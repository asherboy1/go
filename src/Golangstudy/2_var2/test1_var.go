//四种变量声明方式
package main

import "fmt"

// 声明全局变量 仅能使用方法一、二、三
var f float32 = 3.14

// := 与 var2 属于新定义 如果定义好了 就可以直接用=进行赋值

func main() {
	// 方法一 申明一个变量，默认值为0
	var a int
	fmt.Println("a = ", a)

	// 方法二 申明一个变量，初始化一个值
	var b int = 99
	fmt.Println("b = ", b)

	// 方法三 在初始化的时候 可以申请数据类型 通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)

	// 打印变量数据类型
	fmt.Printf("a = %T b = %T\n", a, b)

	// 字符串直接打印与带数据格式打印  %T 为打印数据类型 %s 打印字符串
	var d string = "hello"
	fmt.Println("d = ", d)
	fmt.Printf("d = %s\n", d)

	// 省去 var2 但是 := 只能在函数体内使用
	e := "hello123"
	fmt.Printf("e = %s\n", e)

	// 打印全局变量
	fmt.Println("f = ", f)

	// 申明多个变量1
	var g, h = 1.1, "tony"
	fmt.Println("g = ", g, "h = ", h)

	// 申明多个变量2
	var (
		i int    = 123
		j string = "lili"
	)
	fmt.Println("i = ", i, "j = ", j)

}
