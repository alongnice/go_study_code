package main

import(
	"fmt"
	"reflect"
)

func main(){
	var i int = 10
	var x,y,z int = 1,2,3
	var(
		a int = 10
		b string = "BBB"
		c bool = true
	)

	fmt.Println("基本变量类型")
	fmt.Println(i,reflect.TypeOf(i))
	fmt.Println(x,reflect.TypeOf(x))
	fmt.Println(y,reflect.TypeOf(y))
	fmt.Println(z,reflect.TypeOf(z))
	fmt.Println(a,reflect.TypeOf(a))
	fmt.Println(b,reflect.TypeOf(b))
	fmt.Println(c,reflect.TypeOf(c))
	
	var (
		d = 9.9
		e = "EEE"
		f = false
	)
	fmt.Println("自动类型推导")
	fmt.Println(d,reflect.TypeOf(d))
	fmt.Println(e,reflect.TypeOf(e))
	fmt.Println(f,reflect.TypeOf(f))

}