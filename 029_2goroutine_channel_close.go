package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var stream_sync = make(chan int, 3)
	// 创建一个缓冲

	wg := sync.WaitGroup{} // 创建一个同步等待组
	wg.Add(1)              //计数器+1

	go func() {
		defer wg.Done()
		for {
			<-time.After(time.Duration(50) * time.Microsecond) //50微秒读取一个数
			n, ok := <-stream_sync
			fmt.Println(n, ok)
			if !ok {
				break //channel被关闭,消费完成
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			<-time.After(time.Duration(50) * time.Microsecond)
			stream_sync <- i
		}
		stream_sync <- -1
		close(stream_sync) //生产完成,关闭channel
	}()

	wg.Wait()
}
