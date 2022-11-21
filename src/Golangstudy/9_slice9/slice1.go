package main

import "fmt"

func main() {
	// 动态slice make可以为其设置容量 长度为初始化
	// 长度为6 容量为10  设置长度后相当于初始化了 为0  int 类型  string 无初始化内容
	// cap默认为len
	slice4 := make([]int, 6, 10)

	//len 打印长度 cap 打印容量
	fmt.Println(len(slice4), cap(slice4))

	//相当于可以进行拓展，最大容量为10
	fmt.Println(slice4)

	//向数组中加入了一个 “7”
	slice4 = append(slice4, 7)
	fmt.Println(slice4)

	//向数组中加入了 “8,9,10”
	slice4 = append(slice4, 8, 9, 10)
	fmt.Println(slice4)

	//如果cap已被填满 再往里添值  则 会根据之前所设置的cap扩容 两倍 -> 两倍扩容
	//之前原有的数组 先被复制入新的“扩容后的数组中”，然后新增元素 再被销毁 返回新的slice对象
	//长度》1024时 扩容 扩容为1/4
	slice4 = append(slice4, 11)
	fmt.Println(slice4, cap(slice4)) //cap -> 20

	//截取输出 与python类似
	fmt.Println(slice4[6:])

	//浅拷贝  这里类似于python中浅拷贝 通过传递了地址指针的方式 达到修改slice5 同时影响slice4的效果
	slice5 := slice4[:6]

	slice5[0] = 123

	fmt.Println(slice4)

	//深拷贝  开辟独立内存空间
	slice_tmp := make([]int, 3)
	//将值一一付给slice_tmp
	copy(slice_tmp, slice4[:3])
	slice_tmp[1] = 456
	fmt.Println(slice_tmp)
	fmt.Println(slice4)
}
