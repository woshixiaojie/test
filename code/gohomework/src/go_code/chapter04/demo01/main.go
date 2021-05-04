package main

import "fmt"

func main()  {
	//讲解 / %
	//说明，如果运算的数都是整数，那么除后，去掉小树部分，保留整数部分
	fmt.Println(10 / 4)
	var n1 float64 = 10 / 4
	fmt.Println(n1)

	//保留小数部分，则需要有小数运算
	var n2 float64 = 10.0 / 4
	fmt.Println(n2)

	//演示 % 的作用
	// a % b = a - a / b * b
	fmt.Println("10 % 3 =",10 % 3)
	fmt.Println("10 % -3 =",10 % -3) //10 - 10 / (-3) * (-3)
 	fmt.Println("-10 % 3 =",-10 % 3)
	fmt.Println("-10 % -3 =",-10 % -3)

	//演示 ++ 和 --
	var i int = 10
	i++
	fmt.Println("i =",i)
	i--
	fmt.Println("i =",i)
}
