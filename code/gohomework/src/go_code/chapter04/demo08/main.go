package main

import "fmt"

func test() int {
	return 90
}
//演示Golang中进制
func main()  {
	var i int = 10

	//二进制输出，在Golang中需要使用%b，作为二进制输出
	fmt.Printf("%b \n",i)

	//八进制：0-7，满8进1，以数字0开头
	var j int = 011 // 9
	fmt.Println("j =",j)

	//16进制 0-9及A-F，满16进1，以0x或0X开头表示
	var k int = 0x11 //
	fmt.Println("k =",k)

}

