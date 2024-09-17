package main

import (
	"fmt"
	"math/rand"
	"time"
)

const mostDelayMicrosecond int = 99

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func P(num chan int, delayRate int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(delayRate*TrueRandom(mostDelayMicrosecond)) * time.Microsecond)
		num <- i
		fmt.Println("creat:", i)
	}
	fmt.Println("end of P()")
}
func C(num chan int, delayRate int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(delayRate*TrueRandom(mostDelayMicrosecond)) * time.Microsecond)
		n, ok := <-num
		if ok {
			fmt.Println("closs :", n)
		} else {
			fmt.Println("empty")
		}
	}
	fmt.Println("end of C()")
}

func main() {
	fmt.Println("========case1: Unbuffered========)")
	nums := make(chan int)
	go P(nums, 1)
	go C(nums, 10)
	time.Sleep(2 * time.Second)
	// 无缓冲,只能接受一次吞后接一次吐,依次重复

	fmt.Println("==========case2: buffered=============")
	nums = make(chan int, 20)
	go P(nums, 1)
	go C(nums, 10)
	time.Sleep(2 * time.Second)
	// 有缓冲,可以接受多次吞后吐多次,直到缓冲区满

	fmt.Println("==========case3: buffered but will deadlock===========")
	nums = make(chan int, 9)
	P(nums, 1)
	go C(nums, 10)
	time.Sleep(2 * time.Second)
	// 有缓冲,可以累计
	// 但如果缓冲区满了,会阻塞
	// 所以P和C的执行顺序很重要

}
