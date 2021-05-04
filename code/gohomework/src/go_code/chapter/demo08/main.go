
package main

import "fmt"

//演示Galang字符型使用，使用byte，由单个字节拼接起来的，字符串由字节组成
func main()  {
	//单精度
	var price float32 = 89.21
	fmt.Println("price =",price)

	//双精度
	var num2 float64 = -112344445.11111111111113
	fmt.Println("num2 =",num2)
	//浮点数 = 符号位 + 指数位 + 尾数位
	//尾数部位可能有精度损失
	//64位精度高于32位

	//浮点型使用细节
	//不受操作系统的影响
	//浮点类型默认float64
	var num5 = 1.1
	fmt.Printf("num5 的数据类型是%T\n",num5)

	//十进制数形式 5,12 , .123
	var num6 = 5.12
	var num7 = .123
	fmt.Println("num6 =",num6,"num7 =",num7)

	//科学计数法形式
	num8 := 5.1234567e2//*10的二次方
	num9 := 6.12345678E2//*10的二次方
	num10 := 2222.11e-2// /10的二次方
	fmt.Println("num8 =",num8,"num9 =",num9,"num10",num10)



	//通常情况下使用float64
}



