package main

import (
	"fmt"
	"runtime"
)

// 打印cpu核心数量
func print_num_goroutine() {
	fmt.Println("num_groutine:", runtime.NumCPU())
}

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	// 创建可变参数的匿名函数or，参数为多个channel，返回一个channel，进行实现
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) { //查看channels的长度
		case 0:
			return nil
		case 1:
			return channels[0] //如果只有一个channel，直接返回
		}

		or_done := make(chan interface{}) //运行到这里说明是多个channel，创建一个or_done的channel

		go func() { //创建匿名函数，使用go关键字运行
			defer close(or_done)   //异常捕获，关闭or_done
			switch len(channels) { //重新获取channels长
			case 2: //如果是2
				select {
				case <-channels[0]: //监听多个端口
				case <-channels[1]:
				default:
					select {
					case <-channels[0]:
					case <-channels[1]:
					case <-channels[2]:
					case <-or(append(channels[3:], or_done)...):
					}
				}
			}
		}()

		// close(or_done)
		return or_done
	}

}
