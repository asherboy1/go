package main

import "fmt"

type Human struct {
	// 无效递归
	//SuperMan

	name string
	age  int
	sex  string
}

func (human *Human) gorun() {
	fmt.Println(human.name, "is running")
}

func (human *Human) gowalk() {
	fmt.Println(human.name, "is walking")
}

// 继承类 子承父类

type SuperMan struct {
	Human //直接继承了Human里的属性、方法
	level string
	age   string
}

func (superman *SuperMan) gofighting() {
	fmt.Println(superman.name, "is fighting")
}

type su struct {
}

func main() {
	person := Human{
		name: "tony",
		age:  20,
		sex:  "male",
	}

	person.gorun()

	// 定义子类对象  结构体属性在栈中 方法在堆中
	person1 := SuperMan{
		Human: Human{
			"tom", 20, "male",
		},
		level: "3",
	}

	person1.gorun()
	person1.gofighting()
	fmt.Println(person1)
}
