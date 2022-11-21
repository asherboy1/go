package main

import "fmt"

func main() {
	//定义array [] 中固定长度  不赋值 默认为0
	var array1 [10]int

	// 通过索引遍历数组
	for i := 0; i < 10; i++ {
		fmt.Println(array1[i])
	}
	// 可通过len(),array1.len 获取数组长度
	//for i := 0; i < len(array1); i++ {
	//
	//}

	//直接定义赋值
	array2 := [10]int{1, 2, 3, 4}
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	//for 类似于enumerate 可打印下标  range
	for index, value := range array2 {
		fmt.Println("index=", index, "value=", value)
	}

	fmt.Printf("%T\n", array2) //注意打印格式时 会打印其长度，这代表每个数组为"固定的"  不同长度 表示不同"类型"数组

	//testarray(array1)  无法将长度为10的数组传入长度为5的函数
}

//只能传入长度为5的数组  为值传递
func testarray(a [5]int) {

}
