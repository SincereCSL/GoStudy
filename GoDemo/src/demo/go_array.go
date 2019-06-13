package main

import "fmt"

//申明数组

//var variable_name [SIZE] variable_type
//     变量名               变量类型

//定义数组balance长度为10 类型为float32
var balance [10] float32

//初始化数组：初始化数组中的{}中的元素个数不能大于[]中的数子
var initBalance = [5]float32{100, 2, 30, 40, 50}

//访问数组元素
var salary float32 = balance[9]

//实例演示
func main() {
	var n [10]int /* n 是一个长度为 10 的数组*/
	var i, j int

	/*初始化n数组的元素*/
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /*设置元素i+100*/
	}

	/*输出每个数组元素的值*/
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	/*
	%表示格式化字符串输出
	目前printf支持以下格式的输出，例如：
	printf("%c",a)；输出单个字符。
	printf("%d",a)；输出十进制整数。
	printf("%f",a)；输出十进制浮点数.
	printf("%o",a)；输出八进制数。
	printf("%s",a)；输出字符串。
	printf("%u",a)；输出无符号十进制数。
	printf("%x",a)；输出十六进制数。
	*/
}
