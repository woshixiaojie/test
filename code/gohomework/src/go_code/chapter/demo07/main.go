
package main

import "fmt"

//演示Galang小数类型/浮点类型使用
func main()  {
	var c1  byte = 'a'
	var c2 byte = '0' //字符0

	//直接输出byte，其实是输出对应的字符的码值
	fmt.Println("c1 =",c1,"c2 =",c2)
	//如果我们希望输出对应的字符，需要格式化输出
	fmt.Printf("c1 =%c c2 =%c\n",c1,c2)

	//var c3 byte = '北'
	var c3 int = '北'
	fmt.Printf("c3 =%c c3的编码值=%d",c3,c3)


}



