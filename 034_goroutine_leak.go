package main

import (
	"fmt"
	"time"
)

func main() {
	var do_work = func(strings <-chan string) <-chan interface{} { //定义do_work函数，只接受只读string通道作为输入 返回一个只读的接口
		completed := make(chan interface{}) //创建一个只读的接口通道
		go func() {                         //构建匿名foroutine
			defer fmt.Println("do_work exited") //退出时打印
			defer close(completed)              //退出时关闭接口通道

			for s := range strings { //遍历输入的string通道
				fmt.Println(s) //打印
			}
		}()
		return completed
	}
	ch := make(chan string)
	do_work(ch)
	fmt.Println("Done.")
	close(ch)
	time.Sleep((5 * time.Second))
}
