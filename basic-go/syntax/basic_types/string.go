package main

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	// She said:"hello,go!"
	println("She said:\"hello,go!\"")
	println(`我可以换行
这是新的一行
但是这里不能有反引号`)

	println("Hello," + "GO")
	//字符串只能和字符串拼接
	//println("Hello," + 123)
	//正确的字符串和数字拼接
	println(fmt.Sprintf("hello %d", 123))

	//string 包里放着各种字符串相关操作的方法，需要的时候再查阅

	println(len("你好"))                      // 输出 6
	println(utf8.RuneCountInString("你好"))   // 输出 2
	println(utf8.RuneCountInString("你好go")) // 输出 4
}
