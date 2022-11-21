package main

import "fmt"

func returndefer() int {
	fmt.Println("this is defer7")
	return 0
}

func returnfunc() int {
	fmt.Println("this is func")
	return 0
}

func mixf() int {
	//return 优先级还比 defer7 高，defer为最后的最后

	defer returndefer()
	return returnfunc()
}

func main() {
	mixf()

}
