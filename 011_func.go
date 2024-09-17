package main

import (
	"fmt"
	"strings"
	"reflect"
)

func add(a float64, b float64) (sum float64) {
	sum = a + b
	return sum
}
func sub(a float64, b float64) float64 {
	return a - b
}
func mul() (string, string){
	return "return1","return2"
}

func sum(inputs ...int) (sum int) {
	sum = 0
	for _, v := range inputs {
		sum += v
	}
	return sum
}

func combine2s(strA string, strB string) (strAB string) {
	strAB = strA + strB
	return strAB
}


func main() {
	a, b := 7.0,8.0
	fmt.Println(add(a,b))
	fmt.Println(sub(a,b))
	fmt.Println("===============================")
	
	var arr_str [2]string
	arr_str[0], arr_str[1] = mul()
	slice_form_arrstr := arr_str[0:]
	fmt.Println(slice_form_arrstr)
	fmt.Println("mul return Joined:", strings.Join(slice_form_arrstr, " "))
	fmt.Println("===============================")
	
	// 可变参数用法
	fmt.Println("sum 1~3", sum(1, 2, 3))
	fmt.Println("sum 1~3", sum(1, 2, 3,4, 5, 6))
	vec := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("sum 1~9", sum(vec...))
	fmt.Println("sum type:",reflect.TypeOf(sum))
	
	
	var get_sum func(...int)(int)	// c++ function 对象
	fmt.Println(get_sum)
	get_sum = sum	//c++ bind
	fmt.Println(get_sum(1,2,3,4,5,6))

	fmt.Println("===============================")
	// 指向函数的参考
	addcopy := add
	fmt.Println("addcopy", addcopy(1,2))
	
	
	fmt.Println("===============================")
	// 自定义函数类别 别名
	type two_string_func func(string, string) string
	type aliasof_tsf func(string, string) string
	var loc_join_2str two_string_func = combine2s
	var loc_join_2str_alias aliasof_tsf = combine2s
	fmt.Println(loc_join_2str("hello", "world"))
	fmt.Println(loc_join_2str_alias("hello", "world"))
	fmt.Println("loc_join_2str type:", reflect.TypeOf(loc_join_2str))
	fmt.Println("loc_join_2str_alias type:", reflect.TypeOf(loc_join_2str_alias))
}
