package main

import (
	"fmt"
	"reflect"
)

func main(){
	s1 := make([]int,5);
	s2 := s1;
	fmt.Println(s1,reflect.TypeOf(s1))
	fmt.Println(s2,reflect.TypeOf(s2))

	s2[2]=9
	fmt.Println("s1=",s1)
	fmt.Println("s2=",s2)
	fmt.Println("&s1=%p\t",&s1)
	fmt.Println("&s2=%p\t",&s2)
	// s2是s1的引用

	s3 := []int{3:3, 6:6, 9:9}
	fmt.Println("s3=",s3)
}