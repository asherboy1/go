package main

import "fmt"

// 空接口  万能 通用的接口-数据类型   所有函数、定义都支持interface
func myfunc(arg interface{}) {
	fmt.Println("myfunc is called")
	fmt.Println(arg)

	// interface{} 提供“类型断言“ 机制
	content, assert := arg.(interface{}) // struct string int float interface....
	if !assert {
		fmt.Println("is not string")
	} else {
		fmt.Println("is string")
		fmt.Println(content)
	}
}

type Book struct {
	name string
}

func (book Book) mybook() {
	fmt.Println(book.name)
}
func main() {
	book1 := Book{"zhutou"}
	myfunc(book1) //对象
	// 断言测试
	myfunc(100)
	myfunc("okko")
}
