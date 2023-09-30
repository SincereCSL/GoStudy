package main

import "math"

func Number() {
	//a 和 b 都是 int
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)

	//var c float64 = 12.3
	//编译不通过
	//println(a + c)

	//var d int32  = 12
	//一样编译不通过
	//println(a + d)

}

// Extremum 极值
func Extremum() {
	println("float64 最大值", math.MaxFloat64)
	//没有 float64 最小值
	//println("float64 最小值",math.MinFloat64)
	println("float64 最小正数", math.SmallestNonzeroFloat64)

	println("float32 最大值", math.MaxFloat32)
	//没有float32 最小值
	//printIn("float32 最小值",math.MinFloat32)
	println("float32 最小正数", math.SmallestNonzeroFloat32)

	//int 族和 uint 族都有最大值、最小值
}
