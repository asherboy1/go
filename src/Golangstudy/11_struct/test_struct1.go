package main

import "fmt"

// 声明一种行的数据类型 myint,为int一个别名

type myint int

func main1() {
	var a myint = 999
	fmt.Println("a = ", a)
	fmt.Printf("a = %T\n", a)

	// 引用结构体
	var new_book book
	new_book.title = "Python"

	fmt.Println("title1 = ", new_book.title)
	if &new_book.auth != nil {
		fmt.Println("auth = aa", new_book.auth)
	}

	// 多态
	var new_book2 book
	new_book2.title = "Java"
	new_book2.auth = "Tony"
	fmt.Println("title2 = ", new_book2.title)
	fmt.Println("auth2 = ", new_book2.auth)

	var tmp string
	tmp = new_book2.title
	tmp = "change"
	fmt.Println(tmp, new_book2.title)

	// test1
	new_obj := printbook(new_book2)
	// 值传递 不会改变原有的数据
	fmt.Println(new_book2)
	fmt.Println(new_obj)

	// test2
	new_obj1 := printbook1(&new_book2)
	fmt.Println(new_obj1)
	fmt.Println(new_book2)
}

// test1 传递bookin 无法改变原有的值
func printbook(bookin book) book {
	bookin.auth = "lili"
	bookin.title = "c++"
	return bookin
}

// test2 传递地址 可改变原有的值
func printbook1(bookin *book) string {
	bookin.auth = "lucy"
	bookin.title = "c"
	return "ok"
}

// 定义结构体 复杂 多种的数据类型 融在一起  多态性质 数据属性的一个定义与集合 字典的字段类型集合体
type book struct {
	title string
	auth  string
}
