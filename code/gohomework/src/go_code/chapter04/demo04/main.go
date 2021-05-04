package main

import "fmt"

//声明一个函数（测试）
func test_true() bool {
	fmt.Println("test_t.....")
	return true
}

func test_false() bool {
	fmt.Println("test_f.....")
	return false
}

//逻辑运算，用于连接多个条件（一般来讲就是关系表达式），最终结果也是一个bool值
func main()  {
	// && 也叫短路与，第一个结果为 false，结果为false
	// || 也叫短路或，第一个结果为true，结果为true
	var i int = 10
	//短语与
	if i > 9 && test_true(){
		fmt.Println("ok1")
	}

	if i < 9 && test_true() {
		fmt.Println("ok2")
	}

	//短路或
	if i > 9 || test_true() {
		fmt.Println("ok3")
	}

	//if i < 9 || test_true(){
	//	fmt.Println("ok4")
	//}
	//&&，逻辑与运算符，两边都为true，则为true，否则为false
	//（A && B）为false

	// ||逻辑或运算符，如果两边有一个为True，则为True，否则为False
	//（A || B）为 True

	// ! 逻辑非运算符。如果条件为True，则逻辑为False，否则为True，取反
	//!(A && B) 为 True

	//演示逻辑运算符的使用
	//var age int = 40

	//Golang中 if 不用扩起来
	//if age > 30 && age < 50{
	//	fmt.Println("ok1")
	//}
	//
	//if age > 30 && age < 40{
	//	fmt.Println("ok2")
	//}
	//
	//if age > 30 || age < 50{
	//	fmt.Println("ok3")
	//}
	//
	//if age > 30 || age < 40 {
	//	fmt.Println("ok4")
	//}
	//
	//if !(age > 30 || age < 40) {
	//	fmt.Println("ok5")
	//}
}
