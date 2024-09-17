package main

import "fmt"

func demo_type_assertion(any_type interface{}) {
	switch any_type.(type) {
	case int:
		fmt.Printf("type is int, val = %d\n", any_type.(int))
	case string:
		fmt.Printf("type is string, val = %s\n", any_type.(string))
	default:
		fmt.Println("type is unknown, val = ", any_type)
	}
}

func main() {
	var a interface{}
	var b interface{}
	var c interface{}
	a = 77
	b = "hello"
	c = 3.14

	demo_type_assertion(a)
	demo_type_assertion(b)
	demo_type_assertion(c)

}
