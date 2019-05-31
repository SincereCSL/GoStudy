package main

import "fmt"

func main()  {
	//声明一个变量并初始化
	var a = "go learn"
	fmt.Println(a)

	//没有初始化为零值
	var b int
	fmt.Println(b)

	//bool 零值为false
	var c bool
	fmt.Println(c)
}
