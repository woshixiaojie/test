
package main

import "fmt"

//演示Golang中指针类型
func main()  {
	//基本数据类型，变量存的就是值，也叫值类型

	//基本数据类型在内存中的布局
	var i  int = 10
	//i的地址是多少，&i
	fmt.Println("&i在内存中的地址是",&i)

	//指针类型，指针变量存的是一个地址，这个地址指向的空间才是值
	//比如：var ptr *int = &num
	//指针在内存的布局
	//1、ptr 是一个指针类型
	//2、ptr的类型是 *int
	//3、ptr  本身值是 &i
	var ptr *int = &i
	fmt.Printf("ptr = %v \n",ptr)
	//ptr的地址
	fmt.Println("ptr在内存中的地址",&ptr)

	/*获取指针类型的值，使用 * ，比如 var ptr *int，使用
	*ptr获取ptr指向的值
	 */
	fmt.Printf("ptr 指向的值 %v",*ptr)

	//修改num1的值
	var num1 int
	fmt.Println("num在内存中的地址是",&num1)
	var ptr1 *int = &num1
	*ptr1 = 123
	fmt.Printf("num修改后的值为%v",num1)
}

//值类型和引用类型的使用特点
//值类型：变量直接存储值，内存通常在栈中分配引用，
//引用类型：变量存储的是一个地址，这个地址对应的空间才是真正存储数据值




