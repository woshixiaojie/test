package main

import "fmt"

func test() int {
	return 90
}

func main()  {
	//求两个数的最大值
	var n1 = 100
	var n2 = 200
	var max int
	if n1 > n2 {
		max = n1
	}else {
		max = n2
	}
	fmt.Printf("最大值是%d \n",max)

	//求出三个数的最大值
	//思路：先求出两个值的最大值，然后与第三个值对比，求出三个中的最大值
	var n3 = 300
	if n3 > max {
		max = n3
	}
	fmt.Println("最大值是",n3)
}
