package main

//1.外部导包 注意 外部包的函数名为大写 才能外部使用
//2.开启模块支持后（存在项目下go.mod），并不能与$GOPATH共存,所以把项目从$GOPATH中移出即可   在设置项目包中移除
//3.初始化gomod go mod init 项目名    注意是src的上级目录
//4.注意init执行顺序，当我导入一个包中 存在两个go文件,会按照构筑顺序 执行init
//5.go语言 如果导入 就需使用里面的函数

//#--------------------
//1.匿名导入 如果仅仅想导入包 启用init函数 而不调用其中函数 就需要在 导入前面 加入_    例子： _ "codeForGo/src/Golangstudy/package5/lib1"
//2.用别名 import mylib1 "codeForGo/src/Golangstudy/package5/lib1"  mylib1为别名
//3.用方法导入当前包 import . "codeForGo/src/Golangstudy/package5/lib1"  将所偶方法导入该包 但容易起冲突

import "codeForGo/src/Golangstudy/5_package5/lib1"
import "codeForGo/src/Golangstudy/5_package5/lib2"

func main() {
	lib1.Lib1test()
	lib2.Lib2test()
	lib1.Test()
}
