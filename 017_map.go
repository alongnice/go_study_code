package main

import "fmt"

// ignore 函数接受任意类型的参数，避免未使用变量的报错
func ignore(v ...interface{}) {}

func ListMapInfo(theMap map[string]int) {
	fmt.Println(theMap)
	fmt.Printf("len = %d\n", len(theMap))
}

func main() {
	var table map[string]int = make(map[string]int)

	// add
	table["apple"] = 10
	table["banana"] = 87
	table["Cat"] = 9
	table["Dog"] = 69
	fmt.Println(table)

	// delete
	delete(table, "Dog")

	fmt.Println(table)
	ListMapInfo(table)

	fmt.Println(table["SDCFVGYBHUNJ"]) //0 容易产生误解
	val, ok := table["SDCFVGYBHUNJ"]
	fmt.Println(val, ok)

	mp2 := map[string]string{
		"H": "h",
		"T": "t",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}
	// fmt.Println(mp2)
	ignore(mp2)
}
