package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	mem_cosumed := func() uint64 { //创建一个函数，返回内存使用情况
		runtime.GC()             //强制垃圾回收
		var s runtime.MemStats   //创建一个结构体
		runtime.ReadMemStats(&s) //读取内存使用情况
		return s.Sys             //返回系统内存使用情况
	}

	var c <-chan interface{} //只能取的单向通道
	var wg sync.WaitGroup    //等待组
	var noop = func() {      //空函数
		wg.Done()
		<-c // 未曾放入内容将卡死
	}

	const num_goroutines = 1e4
	wg.Add(num_goroutines)                //添加等待组
	before := mem_cosumed()               //获取内存使用情况
	for i := num_goroutines; i > 0; i-- { //循环创建goroutine,创建很多空函数
		go noop()
	}
	wg.Wait()

	after := mem_cosumed()                                          //获取内存使用情况
	fmt.Printf("%.3fkb", float64(after-before)/num_goroutines/1000) //计算每个goroutine的内存使用情况

}
