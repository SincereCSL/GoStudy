package main

// 全局变量定义
// 首字母大写，全局可以访问
var Global = "全局变量"

// 首字母小写，只能在这个包里使用 其子包也不能使用
var internal = "包变量，私有变量"

var (
	First  string = "1"
	Second int    = 2

	aa string = "test123"
)

// := 只能用于局部变量，全局变量不可以
//test := 123

func main() {
	/// 局部变量定义
	//int 是灰色的，是因为goland 自己可以做类型推断，它认为你可以省略
	var a int = 123
	println(a)

	//这里我们省略了类型
	var b = 234
	println(b)

	//这里的 uint 不能省略，因为省略之后，不加uint类型 这里的456 会被解释为 int 类型
	var c uint = 456
	println(c)

	//这是一句无法编译通过，因为 golang的是强类型语言，并且不会帮你做任何转换
	//println(a == c)

	var (
		d string = "test"
		e int    = 8
	)
	println(d, e)

	//变量声明 :=
	//只能用于局部变量，即方法内部
	// Golang 使用类型推断来推断类型。数字会被理解为 int 或者 float64 (所以其他类型的数字，就得用var来声明)
	f := 123
	println(f)

	//变量声明易错点：
	// 1、变量声明了没有使用
	//var g = 112233
	// 2、类型不匹配
	//var h int = "123"
	//println(h)
	// 3、同作用域下，同名变量只能声明一次
	//var i int = 123
	//var i int = 234
	//println(i)

	//这样不可以，全局变量中声明了 aa
	//var aa int = 123
	//println(aa)
}
