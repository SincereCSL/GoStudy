package main

func Bool() {
	//bool 运算：
	//且（AND）: a && b
	//或 (OR) ： a || b
	//取反 ：!a

	//组合取反是一个很容易犯错的地方：
	// !(a && b) 等价于什么？
	// !(a || b) 等价于什么？
	var a bool = true
	var b bool = false
	//println(!(a && b)) // 输出 true
	//println(!(a || b)) // 输出 false

	var c bool = a || b
	println(c) //true

	var d bool = a && b
	println(d) //false

	var e bool = !a
	println(e) //false

}
