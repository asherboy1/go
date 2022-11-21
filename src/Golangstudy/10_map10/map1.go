package a_10_map10

import "fmt"

func main() {
	//定义map 类似于dict
	//声明map key为string value为string
	//目前仅为声明 为开辟内存空间
	var mymap map[string]string

	//判断是否为未分配内存空间的map
	if mymap == nil {
		fmt.Println("this is empty map")
	}

	//使用map前，需要用make 给map分配空间
	//如果超过所分配空间，底层会自动去增加分配空间
	//map -> hash  内部属于hash排序

	mymap = make(map[string]string, 10)

	mymap["one"] = "tony"
	mymap["two"] = "lili"
	fmt.Println(mymap)

	//声明方式二  动态空间 无需设置大小
	mymap2 := make(map[string]int)
	mymap2["one"] = 1
	mymap2["two"] = 2
	fmt.Println(mymap2)

	//声明方式三  map 中有逗号
	mymap3 := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	fmt.Println(mymap3)
}
