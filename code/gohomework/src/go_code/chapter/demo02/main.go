//变量使用注意事项
//1.变量表示内存中的一个存储区域
//2.该区域有自己的名称（变量名）和类型（数据类型 ）

//Golang变量使用的三种方式
package main

import "fmt"

func main()  {
	//1.制定变量类型，声明后若不负值，使用默认值
	var i int
	//int 默认值是0
	fmt.Println("i=",i)

	//2.根据值自行判定变量类型（类型推到）
	var num = 10.11
	fmt.Println("num = ",num)

	//3.省略var，注意 :=左侧的变量不应该是已经声明过的，否者会导致编译错误
	//等价于 var name string    name = "xiaojie"
	// ：不能省略
	name := "xiaojie"
	fmt.Println("name=",name)
}

//多变量声明
