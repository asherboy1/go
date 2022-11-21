package main

import "fmt"

func main() {
	// 动态数组 []中不设置长度
	myarray := []int{1, 2, 3, 4, 5, 6} //动态数组，切片slice 不指定长度
	testarray1(myarray)

	fmt.Println(myarray[0])
}

//接收传参 也可以[] 动态接收数组  动态数组创建 与 传参 还是遵循一一对应
//[]传递为引用传递 不为值传递
func testarray1(a []int) {
	// _ 表示匿名变量 不会使用到的变量
	for _, value := range a {
		fmt.Println(value)
	}

	//引用传递 可以直接改变传入数组里的值
	a[0] = 100
}
