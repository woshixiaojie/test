
package main

import (
	"fmt"
	"unsafe"
)

//演示Galang中的整数类型的使用
func main()  {
	var i  int = 1
	fmt.Println("i =",i)

	//测试int8的范围，有符号
	//int16,int32 依次类推
	var j int8 = 127
	fmt.Println("j =",j)

	//测试uint8，无符号
	var k uint8 = 255
	fmt.Println("k =",k)

	//rune等价于int32，表示一个Unicode码，byte的使用
	var a int = 89000
	fmt.Println("a =",a)
	var c  byte = 255 //byte等价于uint8，用于存储字符
	fmt.Println("c =",c)

	//整形的使用细节
	var n = 100 //n的数据类型
	fmt.Println("a =",a)
	//查看某个变量的数据类型
	//fmt.Printf()，可用于做格式化输出
	fmt.Printf("n 的类型是%T\n",n)

	//查看变量占用字节大小
	var n2 int32 = 10
	//unsafe.Sizeof() 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 占用的字节数是%d",unsafe.Sizeof(n2))

	//在程序正确运行的情况下，尽量使用占用空间小的原则，占用空间小的数据类型

	//bit计算机中最小的存储单元，byte计算机中基本存储单元
	//1byte = 8bit
}


