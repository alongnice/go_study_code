package main

import "fmt"

type demo_builtln_type struct{}

func (dbit demo_builtln_type) DivideByZero() {
	var zero = 0
	var impossible = 1 / zero
	fmt.Println(impossible)
}
func (dbit demo_builtln_type) Error() string {
	return "the function implements the error interface"
}
func main() {
	var dbit_as_error error = &demo_builtln_type{}
	fmt.Println(dbit_as_error.Error())
}
