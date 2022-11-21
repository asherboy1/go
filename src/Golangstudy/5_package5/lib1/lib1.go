package lib1

import "fmt"

//init函数 相当于开始 setup
func init() {
	fmt.Println("lib1 run!")
}

// 当前lib1包提供的 方法
// 注意！ 导出的函数 首字母必须大写 私有方法 为小写

func Lib1test() {
	fmt.Println("funclib1 run")
}
