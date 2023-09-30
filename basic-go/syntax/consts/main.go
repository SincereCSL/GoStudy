package main

// Go语言中 包名、目录名使用 _ (下划线)命名，变量名、方法名、类型名都是驼峰命名

// 常量声明const
// 首字符是否大写控制了访问性，大写包外可访问
// 驼峰命名
// 支持类型推断
// 无法修改值
const External = "包外可以访问"
const internal = "包内可以访问"

// 支持类型推断
const (
	a = 123
)

// 常量申明 iota 用法
// iota 可以方便地控制常量初始化
// 可以使用 iota,可以使用iota的数学运算，包括位移操作
// 只有主动赋值或者另起一个iota 才会从头开始计算值
const (
	StatusA = iota
	StatusB
	StatusC
	StatusD

	//插入一个主动赋值的就中断了 iota
	StatusE = 6
	StatusF
)

const (
	//DayA = iota << 1
	//DayA = iota >> 1
	//DayA = iota + 1
	DayA = iota*12 + 13
	DayB
	DayC
)

func main() {
	// 局部常量
	//常量用不用都可以编译通过，不报错
	const a = 123

	//常量不能被修改
	//a = 456
}
