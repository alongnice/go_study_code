package main

import "fmt"

type Request struct {
	args        []int
	f           func([]int) int
	result_chan chan int
}

func worker(requests <-chan *Request) {
	for req := range requests {
		result := req.f(req.args)
		req.result_chan <- result
	}
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

// 异步并发模式，通过channel传递request，worker函数处理request，并将结果返回
func main() {
	ans_chan := make(chan *Request)
	// 创建request结构体对象 ans_chan
	go worker(ans_chan)
	// 执行worker函数，将request传递给ans_chan
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// 将request传递给ans_chan
	ans_chan <- request
	// 从request.result_chan中获取结果
	fmt.Println("ans is ", <-request.result_chan)
}
