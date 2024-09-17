package main

import (
	"fmt"
	"time"
)

func main() {
	var put_revover_msg = func() {
		fmt.Println("al`s error: %s \n", recover())
	}

	defer put_revover_msg()

	// 1. defer
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	for i := 0; i < 7; i++ {
		defer fmt.Println("defer ", i)
	}
	fmt.Println("last defer")

	if 0 == time.Now().Second()%2 {
		panic("al`s error")
	} else {
		var zero = 1 - 1
		fmt.Println("zero:", zero)
		var impossible = 1 / zero
		fmt.Println("impossible:", impossible)
	}
}
