package main

import "fmt"

func Byte() {
	var a byte = 'a'
	//这是输出的是 a 的 ASCII 表达 97
	println(a)

	println(fmt.Sprintf("%c", a))

	var str string = "this is string"
	// []byte
	var bs []byte = []byte(str) // 类型转换
	//println(str, bs)
	println(bs)
}
