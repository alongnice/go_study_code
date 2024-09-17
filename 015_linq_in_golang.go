package main

import (
	"fmt"
	// "reflect"
	"strings"

	"github.com/ahmetb/go-linq"
)

func main() {
	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var words []string = []string{"Cat", "Dog", "People", "Cup", "Apple", "Boss", "Zoo", "X-ray", "Github"}

	// linq 基本用法
	fmt.Println(linq.Range(99, 3).Results())                        // 99 100 101
	fmt.Println(linq.From(numbers).SumInts())                       // 0~11 -> sum 66
	fmt.Println(linq.From(numbers).Average())                       // 0~11 -> avg 5.5
	fmt.Println(linq.From(numbers).Min(), linq.From(numbers).Max()) // min max

	fmt.Println(linq.From(words).Last())  // Github
	fmt.Println(linq.From(words).First()) // Cat

	fmt.Println(linq.From([]int{1, 2, 3, 4, 4, 4, 4, 4, 5}).Distinct().Results()...) // set
	fmt.Println(linq.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Reverse().Results()...)  // reverse

	fmt.Println("=====================================================")

	// 大写所有字母
	var only_up_word = make([]string, 0)
	linq.From(words).Select(func(s interface{}) interface{} {
		return strings.ToUpper(s.(string))
	}).ToSlice(&only_up_word)
	fmt.Println(only_up_word)

	// 拼接所有字母
	linq.From(words).ForEach(func(s interface{}) {
		fmt.Print("+", s.(string))
	})
	fmt.Println()

	// 返回boss的所在位置的下标
	index := -1
	linq.From(words).
		SelectIndexed(func(i int, s interface{}) interface{} {
			if s.(string) == "Boss" {
				index = i
			}
			return nil
		}).Results()
	// 输出找到的下标
	if index != -1 {
		fmt.Printf("boss at %d\n", index)
	} else {
		fmt.Println("boss not found")
	}

	// 带o的单词
	var wordsConatinsAlphao = make([]string, 0)
	linq.From(words).Where(func(s interface{}) bool {
		return strings.Contains(s.(string), "o")
	}).ToSlice(&wordsConatinsAlphao)
	fmt.Println(wordsConatinsAlphao)
}
