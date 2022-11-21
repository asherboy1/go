package main

import "fmt"

func main() {
	//1.
	array1 := []int{1, 2, 3, 4, 5} //声明一个slice切片，且初始化数据类型 默认值为1，2，3，4，5 长度为5
	fmt.Println(len(array1))

	//2.
	//该声明方法 只是定义了一个空的 切片  未开辟空间!! 无法赋值
	var slice1 []int

	//开辟空间
	//为slice1 开辟了一个数据类型为int 空间为5  可以利用索引赋值
	slice1 = make([]int, 5)
	fmt.Println(slice1)

	//定义+开辟空间
	var slice2 = make([]int, 3)
	fmt.Println(slice2)

	slice3 := make([]int, 10)
	fmt.Println(slice3)

	//判断slice是否开辟空间  nil
	if slice3 == nil {
		fmt.Println("未开辟空间")
	} else {
		fmt.Println("开辟了空间")
	}

}
