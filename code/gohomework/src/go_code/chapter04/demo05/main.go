package main

import "fmt"

func test() int {
	return 90
}

func main()  {
	//赋值运算符的使用演示
	//var i int
	//i = 1

	//有两个变量，a和b，要求将其进行交换。最终打印结果
	//
	a := 9
	b := 10
	fmt.Printf("a现在的值是%d，b现在的值是%d \n",a,b)
	//定义一个临时变量，先把a的值付给t
	t := a
	a = b
	b = t
	fmt.Printf("a之后的值是%d，b之后的值是%d \n",a,b)

	//赋值运算
	a += 7
	fmt.Printf("%d \n",a)

	//赋值运算符从右向左
	var c int
	c = a + 3
	fmt.Println(c)

	//赋值运算的左边，只能是变量，右边，可以是变量、表达式、常量值
	var d int
	d = a
	d = 8 + 2 * 8// = 右边是表达式
	d = test()// = 右边是函数
	d = 80 //等号右边是常量
	fmt.Println(d)

	//有两个变量，交换，但不允许中间变量
	var a1 = 10
	var b1 = 20

	a1 = a1 + b1
	b1 = a1 - b1 // b1 = a1 + b1 - b1 => b1 = a1
	a1 = a1 - b1// a1 = a1 + b1 - a1 => a1 = b1
	fmt.Printf("交换后a1的值是%d，b1的值是%d",a1,b1)

	//运算符优先级
	/*
	只有单目运算、赋值运算是从右向左

	大致的顺序整理
	1、括号 ++ --
	2、单目运算,对一个值运算，取地址，取值
	3、算数运算 * / % + -
	4、位移运算 << >>
	5、关系运算 < <= > >=
	6、位运算
	7、逻辑运算 && ||
	8、赋值运算 = += -= *= /=
	9、逗号，
	*/
}
