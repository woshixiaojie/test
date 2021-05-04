
package main

import "fmt"


//演示Golang中 + 的演示
func main()  {

	var i = 20
	var j = 30
	var k = i + j//做加法运算
	fmt.Println("k = ",k)

	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2//做拼接操作
	fmt.Println("res =",res)
}


