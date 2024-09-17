package main

import "fmt"

type IFuncAB interface {
	FuncA()
	FuncB()
}
type IFuncA interface {
	FuncA()
}
type Subject struct{}

func (s Subject) FuncA() {
	fmt.Println(("FuncA()"))
}
func (s Subject) FuncB() {
	fmt.Println(("FuncB()"))
}

func main() {
	var objab IFuncAB = &Subject{}
	objab.FuncA()
	var obja IFuncA = objab
	obja.FuncA()

	// var obja IFuncA = &Subject{}
	// obja.funcA()
	// var objab IFuncAB = obja
	// objab.Func()
}
