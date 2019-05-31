package main

//类型相同多个便量，非全局变量

var x, y int
var (
	a int
	b bool
)

var c, d int =1, 2
var e, f = 123, "hello go"

func main()  {
	g, h := 123, "hello"
	println(a,b,c,e,f,g,h)
}
