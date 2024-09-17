package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [10]int{};
	slice := make([]int,10)

	fmt.Println(arr, reflect.TypeOf(arr))	//arr
	fmt.Println(slice, reflect.TypeOf(slice))	//vector

	s1 := []int{0,1,2,3,4}
	s2 := reflect.TypeOf(s1)
	fmt.Println(s1, reflect.TypeOf(s1)) // vector
	fmt.Println(s2, reflect.TypeOf(s2))	// vector

}
