package main

import "fmt"

// 函数 a，b 传参 需定义其传入数据类型，也需定义 返回的数据类型
func foo1(a string, b int) string {
	fmt.Println("a = ", a, "b = ", b)

	c := "tony"

	return c
}

// 函数a，b 传参 定义参数数据类型 如果不匹配传入类型，编译不通过，定义返回数据类型 可定义多个，同样 也许多个返回值进行接收 (int,int)
// 返回值为匿名的
func foo2(a string, b string) (int, int) {
	fmt.Println("a = ", a, "b =", b)

	var num1, num2 = 110, 119
	return num1, num2
}

// 设置返回值为有名的，返回值 “提早”定义好了 名称与类型
func foo3(a int, b float32) (r1 int, r2 string) {
	fmt.Println("a = ", a, "b =", b)

	// 给有名变量 赋值  、 也可以直接使用 作为局部变量 默认值为0
	// r1,r2 作用域为foo3中 局部变量
	r1 = 123
	r2 = "this is r2"
	return r1, r2

}

// 返回值为通过一种数据类型 写法
func foo4(a int, b float32) (r1, r2 int) {
	fmt.Println("a = ", a, "b =", b)

	// 给有名变量 赋值
	r1 = 123
	r2 = 456
	return r1, r2

}

func main() {
	test_foo1 := foo1("lili", 123)
	fmt.Println(test_foo1)

	test_foo2num1, test_foo2num2 := foo2("lili", "199")
	fmt.Println(test_foo2num1, test_foo2num2)

	test_foo3num1, test_foo2num3 := foo3(888, 3.14)
	fmt.Println(test_foo3num1, test_foo2num3)

	test_foo4num1, test_foo4num3 := foo4(888, 3.14)
	fmt.Println(test_foo4num1, test_foo4num3)

	// := 与 var2 属于新定义 如果定义好了 就可以直接用=进行赋值
}
