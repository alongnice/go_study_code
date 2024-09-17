package main

import (
	"fmt"
	"reflect"
)

func main(){
	var intarr [9]int
	intarr[0] = 55
	intarr[1] = 66
	fmt.Println(intarr, reflect.TypeOf(intarr))
	fmt.Println(intarr[0], reflect.TypeOf(intarr[0]))
	fmt.Println(intarr[1], reflect.TypeOf(intarr[1]))
	fmt.Println(intarr[2], reflect.TypeOf(intarr[2]))
	fmt.Println("len(intarr)=",len(intarr))
	fmt.Println("=========================")
	// 未初始化则为0
	
	// slice,切片
	arr1 := []int{0,1,2,3,4}
	fmt.Println(arr1)
	fmt.Println(arr1[0:1])
	fmt.Println(arr1[3:5])
	arr1 = append(arr1, 5,6,7,7,7,)
	fmt.Println(arr1)
	
	fmt.Println("=========================")
	
	// 不定长
	arr2 := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr2,len(arr2))
	fmt.Println("=========================")
	
	// 字串的类型
	fmt.Println(arr2[3:4])
	fmt.Println(reflect.TypeOf(arr2[3:4]))
	fmt.Println(arr2[3])
	fmt.Println(reflect.TypeOf(arr2[3]))
	fmt.Println("=========================")

	// 多维矩阵
	var arr3 [][]int = [][]int{[]int{1,2,3},[]int{4,5,6}}
	fmt.Println(arr3)
	arr3_ := [][]int{[]int{1,2,3},[]int{4,5,6}}
	fmt.Println(arr3_)

	// 遍历多维矩阵
	fmt.Println("-------------style 1----------")
	for i:=0; i<len(arr3); i++{
			fmt.Println(arr3[i])
	}
	fmt.Println("-------------style 2----------")
	for _, element := range arr3 {
		fmt.Println(element)
	}
	fmt.Println("-------------style 3----------")
	for index, element := range arr3 {
		fmt.Println(index, element)
	}

	for i:=0; i<len(arr3); i++{
		for j:=0; j<len(arr3[i]); j++{
			fmt.Print(arr3[i][j])
		}
		fmt.Println()
	}
}