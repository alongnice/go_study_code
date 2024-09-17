package main

import (
	"fmt"
	// "reflect"
	// "strings"
)

type Condition_of_int = func(int) (bool)

func is_odd(number int) bool{
	return 1==number&1
}

func is_even(number int) bool{
	return 0==number&1
}

func Filter(Decider Condition_of_int, numbers []int)(result []int){
	for _, number := range numbers {
		if Decider(number) {
			result = append(result, number)
		}
	}
	return result;
}

func main() {
	var arr = []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(Filter(is_odd, arr))
	fmt.Println(Filter(is_even, arr))

}