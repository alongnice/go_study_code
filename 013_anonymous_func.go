package main

import (
	"fmt"
	"reflect"
	// "strings"
)

type Condition_of_int = func(int) (bool)
func Filter(Decider Condition_of_int, numbers []int)(result []int){
	for _, number := range numbers {
		if Decider(number) {
			result = append(result, number)
		}
	}
	return result;
}

func lambad_func(){
	var func_ref func() = func(){
		fmt.Println("this lamabds func")
	}
	func_ref()
}

func main() {
	var arr = []int{1,2,3,4,5,6,7,8,9}

	fmt.Println(Filter(func(tmp int) bool {return tmp>5}, arr))
	fmt.Println(Filter(func(tmp int) bool {return tmp<=5}, arr))

	fmt.Println("type of lambad_func",reflect.TypeOf(lambad_func))
	lambad_func()
}
