package main

import "fmt"

func main()  {
	//假如还有97天放假，问：还有xx个星期天零多少天
	var week int = 97 / 7
	//% 取余
	var day int = 97 % 7
	fmt.Printf("还有%d个星期零%d天\n",week,day)

	//定义一个变量保存华氏温度，华氏温度转化摄氏温度的公式为：
	//5/9*(华氏温度-100)，请求出华氏温度对应摄氏温度的值
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi-100)
	fmt.Printf("华氏温度%v对应摄氏温度的值为%v",huashi,sheshi)
}
