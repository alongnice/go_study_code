package main

import(
	"fmt"
	// "os"
)

func main(){
	for i:=0; i<10; i++{
		fmt.Print(i)
	}
	fmt.Println() // 换行

	var i int
	for i < 10{
		fmt.Print(i)
		i++;
	}
	fmt.Println() // 换行
	
	var j int
	for {
		fmt.Print(j);
		j++
		if j>=10{
			break
		}
	}
	fmt.Println() // 换行

	var arrstring string = "This is a string"
	for _,c := range arrstring[0:]{
		// fmt.Println(reflect.Typeof(c))==int32
		// 转换为Ascii码
		fmt.Print(c)
	}
	fmt.Println()
}