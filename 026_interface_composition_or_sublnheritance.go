package main

import "fmt"

type Subject struct{}

type IBase interface{ FunBase() }
type IDerivative interface {
	IBase
	FunDerivative()
}

func (s Subject) FunDerivative() {
	fmt.Println("FunDerivative()")
}

// ==== 組合(類繼承)約束 ====
// 註解掉這個函式宣告，就會無法通過編譯
// 出現「*Subject does not implement IDerivative (missing FuncBase method)」
// 證明 strcut 一但實作衍生介面，就同時受到內嵌的基礎介面約束。
// 若要滿足這個約束，就必須聯同內嵌的基礎介面一起實作。
// func (s Subject) FuncBase() {
// 	fmt.Println("FuncBase()")
// }

func main() {

	var ref_derivative IDerivative = &Subject{}
	ref_derivative.FunDerivative()
}
