package main

import "fmt"

func main()  {
	//演示关系运算符的使用
	//结果都是bool类型
	//比较运算符 "=="，不能写成"="
	var n1 int = 8
	var n2 int = 9
	fmt.Println(n1 > n2)
	fmt.Println(n1 == n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	flag := n1 > n2
	fmt.Println(flag)
}
