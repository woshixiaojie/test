
package main

import "fmt"

//在函数外定义的变量就是，全局变量
var n1 = 100
var n2 = 200
var name = "xiaojie"

var (
	n3 = 300
	n4 = 4.44
	name1 = "xiajie"
)
//上面的写法可以一次性声明
func main()  {
	//该案例演示golang如何一次声明多个变量
	//var n1,n2,n3 int
	//fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	//一次性生成多个变量的方式2
	//var n1,name,n3 = 100,"toma",1.11
	//fmt.Println("n1=",n1,"name=",name,"n3=",n3)

	//一次性生成多个变量的方式3，同样可以使用类型推导
	//name,n1,n3 := "xiaojie",11,1.11
	//fmt.Println("name=",name,"n1=",n1,"n3=",n3)

	//输出全局变量
	fmt.Println("name=",name,"n1=",n1,"n2=",n2)
	fmt.Println("n3=",n3,"n4=",n4,"name1=",name1)
}


