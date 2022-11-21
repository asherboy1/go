package main

import "fmt"

// 本质是个指针  【内部构造 有个指针指向当前interface 所指向的具体类型，以及当前类型所包含的函数列表】 父类的指针 全部都需要去继承  抽象类 父类指针指向子类对象
// 意义： 搜索鸭子类型！

type Animalbehavior interface {
	Sleep()
	GetColor() string //类型为该方法其返回的数据类型
	GetType() string
}

// 去【继承】interface 不需要 通过继承方式，仅需去实现即可 自动继承

type Animal_cat struct {
	name  string
	kind  string
	color string
}

func (animal *Animal_cat) Sleep() {
	fmt.Println(animal.name, "is sleeping")
}

func (animal *Animal_cat) GetColor() string {
	return animal.color
}

func (animal *Animal_cat) GetType() string {
	return animal.kind
}

type Animal_dog struct {
	name  string
	kind  string
	color string
}

func (animal *Animal_dog) Sleep() {
	fmt.Println(animal.name, "is sleeping")
}

func (animal *Animal_dog) GetColor() string {
	return animal.color
}

func (animal *Animal_dog) GetType() string {
	return animal.kind
}

// 设置公有方法
func ShowAnimalInfo(animalbehavior Animalbehavior) {
	fmt.Println("type:", animalbehavior.GetType())
	fmt.Println("color:", animalbehavior.GetColor())
}

func main() {
	var animal Animalbehavior //接口的数据类型，父类指针
	animal = &Animal_cat{
		name:  "mimi",
		kind:  "dada",
		color: "blue",
	}
	animal.Sleep() //即调用cat的 结构体的方法

	animal = &Animal_dog{
		name:  "zhuangzhuang",
		kind:  "xiaoxiao",
		color: "red",
	}
	fmt.Println(animal.GetColor()) //调用dog的 结构体方法
	ShowAnimalInfo(animal)         //进行调用 此为一种调用方式 也可以传入具体对象的地址

	// 方法2
	cat1 := Animal_cat{
		name:  "nini",
		kind:  "qiqi",
		color: "white",
	}

	dog1 := Animal_dog{
		name:  "bullen",
		kind:  "cowdog",
		color: "black",
	}

	ShowAnimalInfo(&cat1)
	fmt.Println("------------------")
	ShowAnimalInfo(&dog1)

}
