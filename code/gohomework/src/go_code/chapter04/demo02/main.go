package main

import "fmt"

//算数运算符，使用注意事项
func main()  {
	//在Golang中，++ 和 -- 只能独立使用
	//var i int = 8
	//var a int

	//i++,i-- 只能独立使用
	//a = i++
	//a = i--
	//if i++ > 0 {
	//	fmt.Println("ok")
	//}
	var i int= 1
	i++
	i--
	//在Golang中没有 前++，前--
	//++i
	//--i
	fmt.Println(i)
}
