package main

import (
	"fmt"
	"sort"
	"time"

	"math/rand"
)

// func ignore(v ...interface{}) {}

type china_num struct{ val string }
type china_nums []china_num

func (p china_nums) Len() int { return len(p) }
func (p china_nums) Less(i, j int) bool {
	var china_num_tab = map[string]int{
		"零": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9, "十": 10,
	}
	return china_num_tab[p[i].val] < china_num_tab[p[j].val]
}
func (p china_nums) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	var slice_china_num = []china_num{
		{val: "零"},
		{val: "一"},
		{val: "二"},
		{val: "三"},
		{val: "四"},
		{val: "五"},
		{val: "六"},
		{val: "七"},
		{val: "八"},
		{val: "九"},
	}
	fmt.Println("=========init============")
	fmt.Println(slice_china_num)

	rand.Seed(int64(time.Now().Second()))               //生成随机数
	rand.Shuffle(len(slice_china_num), func(i, j int) { //打乱
		slice_china_num[i], slice_china_num[j] = slice_china_num[j], slice_china_num[i]
	})
	fmt.Println("=========shuffle============")
	fmt.Println(slice_china_num)

	sort.Sort(china_nums(slice_china_num))
	fmt.Println("=========sort============")
	fmt.Println(slice_china_num)
}
