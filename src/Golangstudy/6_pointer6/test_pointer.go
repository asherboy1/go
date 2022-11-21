package main

import "fmt"

func changeval(a int) {
	a = 0
	fmt.Println("a = ", a)

}

func main() {
	//值传递
	var a int = 10
	changeval(a)
	fmt.Println(a)
}
