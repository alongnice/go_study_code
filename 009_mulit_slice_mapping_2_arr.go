package main

import (
	"fmt"
	// "reflect"
)

func main() {
	var only_one_arr = [12]int{0,1,2,3,4,5,6,7,8,9,10,11}

	var s0to3 = only_one_arr[:4]
	var s4to7 = only_one_arr[4:7]
	var s9to11 = only_one_arr[9:]

	fmt.Println(only_one_arr)
	fmt.Println(s0to3)
	fmt.Println(s4to7)
	fmt.Println(s9to11)

	var temp_arr = [len(only_one_arr)]int = only_one_arr
	for i:=0; i<len(only_one_arr); i++{
		only_one_arr[i] = temp_arr[len(only_one_arr)-1-i]
	}
	fmt.Println("数组反转")

	fmt.Println(only_one_arr)
	fmt.Println(s0to3)
	fmt.Println(s4to7)
	fmt.Println(s9to11)
}
