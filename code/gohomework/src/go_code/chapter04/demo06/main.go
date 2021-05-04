package main

import "fmt"

func test() int {
	return 90
}

func main()  {
	//演示 & 和 * 的使用
	a := 100
	fmt.Println("a 的地址=",&a)

	//先使用& 把值的地址交给ptr
	var ptr  = &a
	//然后通过 * 取出地址的值
	fmt.Println("ptr 指向的值是=",*ptr)

	//GoLang设计理念
	//一种事情有且只有一种方法完成
	//不支持三元运算法，直接使用 if，else
	var n int
	var i int = 10
	var j int = 20
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Printf("n = %d",n)
}
