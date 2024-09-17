package main

import (
	"fmt"
	"reflect"
	// "strings"
)

func main() {
	var nums []int = []int{0,1,2,3,4,5,6,7,8,9}

	var ref2print_int []func() = nil;
	for _,num:=range nums{
		ref2print_int = append(ref2print_int, func(){
			loc_num := num
			fmt.Println(loc_num)
		})
		ref2print_int[len(ref2print_int)-1]()
	}
	fmt.Println(reflect.TypeOf(ref2print_int))
	for _,f := range ref2print_int{
		f()
	}
}

