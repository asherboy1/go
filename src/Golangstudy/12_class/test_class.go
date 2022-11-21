package main

import (
	"fmt"
)

// 导入中当前go包与导入的 仅能存在一个main 入口函数
//https://blog.csdn.net/lvyibin890/article/details/83651170
// *"codeForGo/src/Golangstudy/12_class/test_get"*/

// 细节 与方法导包一致，如果首字母定义为大写 则 其他包引用时可以访问 该结构体
type people struct {
	// 如果说类的属性首字母大写，表示该属性能够进行访问，否则只能够类的内部访问
	name string
	age  int
	sex  string
}

func (People people) getname() {
	fmt.Println(People.name)
}

// 副本 不影响原有值
func (People people) setname(name string) {
	People.name = name
}

func (People *people) changename(name string) {
	People.name = name
}

func main() {
	// 创建对象
	people1 := people{
		name: "tony",
		age:  19,
		sex:  "male",
	}
	people1.getname()

	// 这里传值还是为值传递 想要改变结构体的值 需要改变原有的值 需用到指针
	people1.setname("tom")
	people1.getname()

	// 传入指针 对于类的方法 默认调用传入的就为该对象的地址  所以不需要&传入地址
	// 注意 对象实例化的方法 与 公共方法传入指针的不同
	people1.changename("lucy")
	people1.getname()
}
