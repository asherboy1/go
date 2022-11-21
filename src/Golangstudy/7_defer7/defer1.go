package main

import "fmt"

func func1() {
	fmt.Println("this is func1")
}

func func2() {
	fmt.Println("this is func2")
}

func func3() {
	fmt.Println("this is func3")
}

func main() {
	//注意压栈顺序 先入后出
	defer func1()
	defer func2()
	defer func3()
}
