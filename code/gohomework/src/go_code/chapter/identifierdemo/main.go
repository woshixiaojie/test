
package main

//演示Golang中标识符的使用，也就是变量
func main()  {
	//不能以数字开头，不能含有空格，-
	// int，float可以使用，但是强烈要求不去使用

	//Golang中严格区分大小写

	//标识符不能有空格

	//_ 是空标识符，用于占位，不能放在开头

	//不能以系统保留关键字作为标识符

	/*
	1、包名尽量和文件夹名字保持一致，采用有简短，意义的名，
	不要和标准库冲突 fmt

	2、变量名、函数名、常量名：采用驼峰法
	var stuName string = "tom" xxxYzz
	var goodPrice float64 = 12345.12345

	3、如果变量名、函数名、常量名首字母大写，则可以被其他的包访问
	；如果首字母小写，则只能在本包中使用
	（首字母大写公有，首字母小写私用）
	 */

}




