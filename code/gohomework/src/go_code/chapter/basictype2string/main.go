
package main

import (
	"fmt"
	"strconv"
)

//演示Golang中基本数据联系转成string使用
func main()  {
	var num1 int = 99
	var num2 float64 = 10.11
	var b bool = true
	//var myChar byte = 'h'
	var str string //空字符串

	//使用第一种方式fmt.Sprintf方法

	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T , str = %q \n",str,str)

	//第二种方式strconv
	str = strconv.FormatInt(int64(num1),10)
	fmt.Printf("str type %T , str = %q \n",str,str)

	//f表示一种格式
	//10：表示表示小数后10位
	//64：表示小数是float64类型
	str = strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("str type %T , str = %q \n",str,str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type %T , str = %q \n",str,str)

	//strconv中Itoa可以直接将int类型转化成str
	str = strconv.Itoa(num1)
	fmt.Printf("str type %T , str = %q \n",str,str)

}



