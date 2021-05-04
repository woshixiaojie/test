
package main

import (
	"fmt"
	"strconv"
)

//演示Golang中string转化成基本数据类型
func main()  {
	var str string =  "true"
	var b bool
	//strconv.ParseBool(str)会返回两个值（value bool，err error）
	//只想获取到value bool ，不要获取err，使用 _ 忽略错误
	b ,_ = strconv.ParseBool(str)
	fmt.Printf("b typr %T b = %v \n",b,b)

	var str2 string = "1233456"
	var n1 int64
	var n2 int
	n1,_ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)
	fmt.Printf("n1 typr %T b = %v \n",n1,n1)
	fmt.Printf("n2 typr %T b = %v \n",n2,n2)

	//返回什么，必须要对应的类型接受
	var str3 string = "123.123"
	var f1 float64
	f1,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("f1 typr %T b = %v \n",f1,f1)

	/*再将String类型转化成基本数据类型时，要确保String类型能够
		转化成一个有效的数据
	 */
	//演示
	var str4 string = "hello"
	var int3 int64 = 11
	int3,_ = strconv.ParseInt(str4,10,64)
	//默认值int3为0
	//int float => 0 , bool => false
	fmt.Printf("int3 type %T int3 = %v \n",int3,int3 )
}




