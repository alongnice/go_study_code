package main

import (
	"fmt"
	// "reflect"
)

func main(){
	const(
		nine = 9.9/1.1
		const_str string = "read only string"
		falseConst bool = false
	)

	fmt.Println(nine)
	fmt.Println(const_str)

	const(
		c0 = iota	// 类似enum
		c1
		c2
	)
	fmt.Println(c0,c1,c2)

	const(
		d0 = iota+1
		d1
		d2
	)
	fmt.Println(d0,d1,d2)

	const(
		x0=9
		x1
		x2
	)
	fmt.Println(x0,x1,x2)
}