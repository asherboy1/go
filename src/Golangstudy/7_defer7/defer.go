package main

import "fmt"

func main() {
	//写defer关键字
	//defer表示在函数体内 结束要去触发的机制 类似与finally finalize 回收对象之前操作
	//可以在函数定义多个defer 依照顺序执行defer
	//defer7 可以执行另一个function 遵循栈规则 先入后出  详见defer1

	defer fmt.Println("main end")

	fmt.Println("main::hello1")
}
