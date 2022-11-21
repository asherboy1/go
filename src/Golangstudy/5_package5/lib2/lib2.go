package lib2

import "fmt"

func init() {
	fmt.Println("lib2 run!")
}

// 当前lib2包提供的 方法
// 注意！ 导出的函数 首字母必须大写 私有方法 为小写

func Lib2test() {
	fmt.Println("funclib2 run")
}
